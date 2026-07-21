import uuid
import datetime
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession
from app.infrastructure.database import engine
from app.infrastructure.models.guardrail_entities import (
    Base,
    PolicyDb,
    SafetyEventDb,
    ViolationDb,
    PiiEventDb,
    SecurityEventDb,
    HumanReviewDb,
)
from app.domain.guardrails import PolicyDto, SafetyEventDto


class GuardrailRepository:
    """Repository managing Database persistence for Safety Policies, Violations, PII/Security audits, and human reviews."""

    async def initialize_tables(self) -> None:
        """Create tables in target database context and seed default governance policies on startup."""
        async with engine.begin() as conn:
            await conn.run_sync(Base.metadata.create_all)

        # Seed default policies using a direct connection
        async with AsyncSession(engine) as session:
            await self._seed_default_policies(session)

    async def _seed_default_policies(self, session: AsyncSession) -> None:
        defaults = [
            ("PII Data Leakage Protection", "pii_leakage", "Redacts name, email, Aadhaar, PAN, SSN, and other identifiers.", "BLOCK"),
            ("Credential & Secret Exposure", "secret_exposure", "Flags and redacts AWS keys, private tokens, passwords, and private certs.", "BLOCK"),
            ("Toxicity & Hate Moderation", "harmful_toxicity", "Identifies hate speech, harassment, physical threat instructions, or self-harm.", "BLOCK"),
            ("Indirect Prompt Injection", "prompt_injection", "Blocks systemic prompt manipulations or jailbreak system templates.", "BLOCK"),
            ("RAG Grounding Hallucinations", "hallucination_grounding", "Computes context alignment citation confidence to target hallucinations.", "WARN"),
        ]
        
        for name, code, desc, action in defaults:
            stmt = select(PolicyDb).where(PolicyDb.code == code)
            res = await session.execute(stmt)
            if not res.scalar_one_or_none():
                db_policy = PolicyDb(
                    id=str(uuid.uuid4()),
                    name=name,
                    code=code,
                    description=desc,
                    action=action,
                    status="ACTIVE",
                )
                session.add(db_policy)
        await session.commit()

    async def get_policies(self, session: AsyncSession) -> list[PolicyDto]:
        stmt = select(PolicyDb)
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            PolicyDto(
                id=r.id,
                name=r.name,
                code=r.code,
                description=r.description,
                action=r.action,
                status=r.status,
            )
            for r in rows
        ]

    async def get_policy_by_code(self, session: AsyncSession, code: str) -> PolicyDb | None:
        stmt = select(PolicyDb).where(PolicyDb.code == code)
        res = await session.execute(stmt)
        return res.scalar_one_or_none()

    async def create_policy(self, session: AsyncSession, name: str, code: str, description: str | None, action: str) -> PolicyDto:
        db_policy = PolicyDb(
            id=str(uuid.uuid4()),
            name=name,
            code=code,
            description=description,
            action=action,
            status="ACTIVE",
        )
        session.add(db_policy)
        await session.commit()
        return PolicyDto(
            id=db_policy.id,
            name=db_policy.name,
            code=db_policy.code,
            description=db_policy.description,
            action=db_policy.action,
            status=db_policy.status,
        )

    async def update_policy(self, session: AsyncSession, code: str, action: str, status: str) -> PolicyDto | None:
        stmt = select(PolicyDb).where(PolicyDb.code == code)
        res = await session.execute(stmt)
        db_policy = res.scalar_one_or_none()
        if not db_policy:
            return None
        
        db_policy.action = action
        db_policy.status = status
        await session.commit()
        
        return PolicyDto(
            id=db_policy.id,
            name=db_policy.name,
            code=db_policy.code,
            description=db_policy.description,
            action=db_policy.action,
            status=db_policy.status,
        )

    async def log_safety_event(
        self, session: AsyncSession, user_id: str, tenant_id: str, input_text: str, output_text: str | None, action_taken: str
    ) -> str:
        event_id = str(uuid.uuid4())
        db_event = SafetyEventDb(
            id=event_id,
            user_id=user_id,
            tenant_id=tenant_id,
            input_text=input_text,
            output_text=output_text,
            action_taken=action_taken,
        )
        session.add(db_event)
        await session.commit()
        return event_id

    async def add_violation(self, session: AsyncSession, event_id: str, policy_id: str | None, rule: str, severity: str, details: str) -> None:
        db_viol = ViolationDb(
            id=str(uuid.uuid4()),
            event_id=event_id,
            policy_id=policy_id,
            rule_triggered=rule,
            severity=severity,
            details=details,
        )
        session.add(db_viol)
        await session.commit()

    async def add_pii_event(self, session: AsyncSession, event_id: str, pii_type: str, masked: str) -> None:
        db_pii = PiiEventDb(
            id=str(uuid.uuid4()),
            event_id=event_id,
            pii_type=pii_type,
            masked_value=masked,
        )
        session.add(db_pii)
        await session.commit()

    async def add_security_event(self, session: AsyncSession, event_id: str, sec_type: str, secret: str) -> None:
        db_sec = SecurityEventDb(
            id=str(uuid.uuid4()),
            event_id=event_id,
            security_type=sec_type,
            secret_detected=secret,
        )
        session.add(db_sec)
        await session.commit()

    async def create_human_review(self, session: AsyncSession, event_id: str) -> str:
        review_id = str(uuid.uuid4())
        db_review = HumanReviewDb(
            id=review_id,
            event_id=event_id,
            status="PENDING",
        )
        session.add(db_review)
        await session.commit()
        return review_id

    async def resolve_human_review(self, session: AsyncSession, event_id: str, approved: bool, reviewer_id: str, notes: str | None) -> None:
        stmt = select(HumanReviewDb).where(HumanReviewDb.event_id == event_id).where(HumanReviewDb.status == "PENDING")
        res = await session.execute(stmt)
        db_rev = res.scalar_one_or_none()
        if db_rev:
            db_rev.status = "APPROVED" if approved else "REJECTED"
            db_rev.reviewer_id = reviewer_id
            db_rev.override_approved = approved
            db_rev.notes = notes
            
            # Update safety event status accordingly
            evt_stmt = select(SafetyEventDb).where(SafetyEventDb.id == event_id)
            evt_res = await session.execute(evt_stmt)
            db_evt = evt_res.scalar_one_or_none()
            if db_evt:
                db_evt.action_taken = "ALLOWED" if approved else "BLOCKED"
                
            await session.commit()

    async def get_safety_events(self, session: AsyncSession, limit: int = 50) -> list[SafetyEventDto]:
        stmt = select(SafetyEventDb).order_by(SafetyEventDb.timestamp.desc()).limit(limit)
        res = await session.execute(stmt)
        rows = res.scalars().all()
        return [
            SafetyEventDto(
                id=r.id,
                user_id=r.user_id,
                tenant_id=r.tenant_id,
                input_text=r.input_text,
                output_text=r.output_text,
                action_taken=r.action_taken,
                timestamp=str(r.timestamp),
            )
            for r in rows
        ]


guardrail_repository = GuardrailRepository()
