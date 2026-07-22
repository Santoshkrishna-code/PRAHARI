import React, { useState } from 'react';
import {
  LayoutDashboard, Activity, BarChart3, FileText, Radio, Box, Package, Wrench,
  Shield, AlertTriangle, ClipboardCheck, Target, Brain, Cpu, Eye, Server,
  Settings, ChevronRight, ChevronDown, Search, Bell, User, Zap, ArrowUpRight,
  ArrowDownRight, Minus, Clock, CheckCircle2, XCircle, AlertCircle, Sparkles,
  MessageSquare, Play, Circle, TrendingUp, ChevronLeft, Command, Filter,
  Download, Plus, RefreshCw, Maximize2, Layers, Thermometer, Gauge, BarChart,
  PieChart, Calendar, MapPin, Hash, ArrowRight, ExternalLink, Wifi, Database,
  Lock, Unlock, FileSearch, GitBranch, MoreHorizontal, LogOut, UserCheck,
  SlidersHorizontal, Rewind, SkipForward, Users, HelpCircle, List, Grid3X3,
  Check, FileCheck, Award, FileWarning
} from 'lucide-react';

import { TelemetryPoint, HealthStatus, Asset, Incident } from './types';
import { UserSession } from './types/auth';
import { Toolbar, ToolBtn, ToolSep, Chart, Metric } from './components/common/CommonUI';

// ═══════════════════════════════════════════════════════════
// WORKSPACE 1: COMMAND CENTER
// ═══════════════════════════════════════════════════════════
export const CommandCenter: React.FC<{
  tele: TelemetryPoint[];
  onReportIncident?: () => void;
  onNavigateToAsset?: (assetId: string) => void;
  onNavigateToPage?: (page: any) => void;
  onGenerateWorkOrder?: (asset: string, desc: string) => void;
}> = ({ tele, onReportIncident, onNavigateToAsset, onNavigateToPage, onGenerateWorkOrder }) => {
  const l = tele[tele.length - 1] || { vib: 11.8, temp: 94.1, psi: 242, kw: 330, flow: 84, t: '09:47' };
  return (
    <div className="h-full flex flex-col bg-[#09090b]">
      <Toolbar>
        <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">COMMAND CENTER</span>
        <ToolSep />
        <span className="text-[10px] text-zinc-500">PLANT ALPHA REFINERY • GULF COAST SITE</span>
        <div className="flex-1" />
        <ToolBtn onClick={onReportIncident} className="!bg-red-600/80 !text-white hover:!bg-red-500">
          <AlertTriangle size={12} /> Report Incident
        </ToolBtn>
        <ToolBtn onClick={() => onNavigateToPage?.('ai-command')}><Brain size={12} /> Open AI Command</ToolBtn>
        <ToolBtn onClick={() => onNavigateToAsset?.('pump')}><Layers size={12} /> Focus Digital Twin</ToolBtn>
      </Toolbar>

      <div className="flex-1 overflow-y-auto p-5 space-y-5">
        <div className="flex items-center justify-between p-4 rounded-xl bg-white/[0.02] border border-white/[0.04]">
          <div>
            <div className="flex items-center gap-2">
              <h2 className="text-base font-bold text-white">Good Afternoon — Plant Alpha Refinery</h2>
              <span className="text-[10px] px-2 py-0.5 rounded-full bg-emerald-500/10 text-emerald-400 border border-emerald-500/20 font-semibold">
                Shift B (143 Operators Online)
              </span>
            </div>
            <p className="text-xs text-zinc-400 mt-1">
              Overall Safety Index: <span className="text-emerald-400 font-bold">94.2/100</span> (↑ +2.1% from yesterday) • 1,284 signals monitored across 47 assets.
            </p>
          </div>
          <div className="text-right text-xs text-zinc-500">
            <div>Last Telemetry Stream: <strong className="text-zinc-300">{l.t}</strong></div>
            <div>AWS Region: <strong className="text-zinc-300">us-east-1</strong></div>
          </div>
        </div>

        <div className="grid grid-cols-6 gap-3">
          {[
            { l: 'Safety Index', v: '94.2', u: '/100', c: 'text-emerald-400', sub: '+2.1% vs yesterday', onClick: () => onNavigateToPage?.('ops-intelligence') },
            { l: 'Incident Rate (TRIR)', v: '0.18', u: '', c: 'text-emerald-400', sub: 'Zero OSHA recordables', onClick: () => onNavigateToPage?.('incidents') },
            { l: 'Active Risks', v: '2', u: 'High', c: 'text-amber-400', sub: 'DC-101 Loop', onClick: () => onNavigateToPage?.('risk') },
            { l: 'Inspection Pass', v: '98.4', u: '%', c: 'text-emerald-400', sub: '122 audited', onClick: () => onNavigateToPage?.('inspections') },
            { l: 'Asset Health', v: '91.2', u: '%', c: 'text-white', sub: '47 monitored', onClick: () => onNavigateToPage?.('assets') },
            { l: 'Permit Compliance', v: '100', u: '%', c: 'text-emerald-400', sub: '28 active LOTO', onClick: () => onNavigateToPage?.('permits') },
          ].map(k => (
            <div key={k.l} onClick={k.onClick} className="p-3.5 rounded-xl bg-white/[0.02] hover:bg-white/[0.04] border border-white/[0.04] cursor-pointer transition-colors">
              <p className="text-[9px] text-zinc-500 uppercase tracking-wider mb-1 font-semibold">{k.l}</p>
              <p className={`text-xl font-bold ${k.c}`}>{k.v}<span className="text-xs text-zinc-500 font-normal ml-0.5">{k.u}</span></p>
              <p className="text-[10px] text-zinc-600 mt-1">{k.sub}</p>
            </div>
          ))}
        </div>

        <div className="grid grid-cols-3 gap-5">
          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <div className="flex items-center gap-2">
                <Sparkles size={14} className="text-indigo-400" />
                <span className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider">AI Supervisor Actions</span>
              </div>
              <span className="text-[10px] text-zinc-600">3 Pending</span>
            </div>
            {[
              { pri: 'high', title: 'Replace bearing assembly — Pump P-102', summary: 'Vibration velocity (11.8 mm/s) trending toward 14.0 mm/s limit. RUL: 18 days. Bearing wear: 72%.', agent: 'Maintenance Agent', confidence: 94, action: 'Generate WO', onAct: () => onGenerateWorkOrder?.('PUMP-P102', 'Bearing race replacement') },
              { pri: 'medium', title: 'Audit Zone B contractor badges', summary: 'Contractor C-4412 detected with expired badge near reactor. Gate B access auto-revoked.', agent: 'Contractor Agent', confidence: 92, action: 'Inspect Vision', onAct: () => onNavigateToPage?.('vision-intel') },
              { pri: 'low', title: 'Schedule lubrication — Compressor C-03', summary: 'PM interval threshold reached in 6 days. Auto-escalation configured in CMMS.', agent: 'Maintenance Agent', confidence: 96, action: 'View Asset', onAct: () => onNavigateToAsset?.('comp') },
            ].map((r, i) => (
              <div key={i} className="p-3.5 rounded-xl bg-white/[0.02] hover:bg-white/[0.04] border border-white/[0.04] transition-colors group">
                <div className="flex items-start gap-2.5">
                  <span className={`w-1.5 h-1.5 rounded-full mt-2 shrink-0 ${r.pri === 'high' ? 'bg-amber-400' : r.pri === 'medium' ? 'bg-indigo-400' : 'bg-zinc-600'}`} />
                  <div className="flex-1 min-w-0">
                    <p className="text-[13px] font-medium text-zinc-200">{r.title}</p>
                    <p className="text-[12px] text-zinc-500 mt-1 leading-relaxed">{r.summary}</p>
                    <div className="flex items-center justify-between mt-2 pt-2 border-t border-white/[0.04]">
                      <div className="flex items-center gap-2 text-[10px] text-zinc-600">
                        <span className="flex items-center gap-1"><Brain size={10} />{r.agent}</span>
                        <span>{r.confidence}% conf.</span>
                      </div>
                      <button onClick={r.onAct} className="px-2 py-1 rounded bg-indigo-600 hover:bg-indigo-500 text-white font-bold text-[10px]">
                        {r.action}
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            ))}
          </div>

          <div className="space-y-3">
            <span className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider flex items-center gap-2">
              <Gauge size={14} className="text-zinc-500" /> Live Parameter Telemetry
            </span>
            {[
              { label: 'Vibration (P-102)', value: l.vib, unit: 'mm/s', max: 16, color: l.vib > 12 ? '#f59e0b' : '#22c55e', note: 'Threshold: 14.0 mm/s' },
              { label: 'Bearing Temperature', value: l.temp, unit: '°C', max: 120, color: l.temp > 95 ? '#f59e0b' : '#818cf8', note: 'Threshold: 100.0°C' },
              { label: 'Line Pressure', value: l.psi, unit: 'PSI', max: 300, color: '#22d3ee', note: 'Nominal 200–260' },
              { label: 'Compressor Power', value: l.kw, unit: 'kW', max: 400, color: '#a78bfa', note: 'Peak: 370 kW' },
              { label: 'Flow Rate', value: l.flow, unit: 'L/m', max: 110, color: '#34d399', note: 'Target: 80 L/m' },
            ].map(g => (
              <div key={g.label} className="p-3 rounded-xl bg-white/[0.02] border border-white/[0.04]">
                <div className="flex justify-between items-baseline mb-1">
                  <span className="text-[11px] text-zinc-400 font-medium">{g.label}</span>
                  <span className="text-sm font-bold text-white">{g.value}<span className="text-[9px] text-zinc-500 ml-0.5">{g.unit}</span></span>
                </div>
                <div className="w-full h-1.5 rounded-full bg-zinc-800 mb-1">
                  <div className="h-1.5 rounded-full transition-all duration-500" style={{ width: `${(+g.value / g.max) * 100}%`, backgroundColor: g.color }} />
                </div>
                <p className="text-[9px] text-zinc-600 text-right">{g.note}</p>
              </div>
            ))}
          </div>

          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <span className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider flex items-center gap-2">
                <Activity size={14} className="text-zinc-500" /> Plant Activity Stream
              </span>
              <div className="flex items-center gap-1">
                <span className="w-1.5 h-1.5 rounded-full bg-emerald-400 animate-pulse" />
                <span className="text-[10px] text-emerald-400 font-medium">Live</span>
              </div>
            </div>
            <div className="space-y-1.5 max-h-[380px] overflow-y-auto pr-1">
              {[
                { t: '09:47:22', src: 'AI Supervisor', msg: 'Dispatched Risk Agent for P-102 vibration anomaly', sev: 'ai' },
                { t: '09:47:18', src: 'Edge Vision', msg: 'CAM-002 PPE violation — no helmet in Zone B', sev: 'err' },
                { t: '09:47:04', src: 'Maintenance', msg: 'RUL recalculated for P-102: 18 days remaining', sev: 'warn' },
                { t: '09:46:51', src: 'Permit Agent', msg: 'PTW-8902 isolation verified on Valve V-88', sev: 'ok' },
                { t: '09:46:38', src: 'Contractor', msg: 'Badge C-4412 expired — access revoked automatically', sev: 'warn' },
                { t: '09:46:22', src: 'Telemetry', msg: 'Vibration probe VP-102 reading 11.8 mm/s', sev: 'warn' },
                { t: '09:46:11', src: 'Edge Vision', msg: 'CAM-001 PPE scan complete — 3 workers verified', sev: 'ok' },
                { t: '09:45:58', src: 'Telemetry', msg: 'Temperature TS-102 nominal at 88°C', sev: 'ok' },
              ].map((e, i) => (
                <div key={i} className="flex items-start gap-2 py-2 px-2.5 rounded-lg bg-white/[0.015] hover:bg-white/[0.03] transition-colors text-[12px] cursor-pointer group">
                  <span className={`w-1.5 h-1.5 rounded-full mt-1.5 shrink-0 ${e.sev === 'err' ? 'bg-red-400' : e.sev === 'warn' ? 'bg-amber-400' : e.sev === 'ai' ? 'bg-indigo-400' : 'bg-emerald-400/60'}`} />
                  <div className="flex-1 min-w-0">
                    <span className="text-zinc-400 group-hover:text-zinc-200 transition-colors">{e.msg}</span>
                    <div className="flex gap-2 mt-0.5 text-[10px] text-zinc-600">
                      <span>{e.src}</span><span>{e.t}</span>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 2: OPERATIONS INTELLIGENCE
// ═══════════════════════════════════════════════════════════
export const OpsIntelligence: React.FC<{
  tele?: TelemetryPoint[];
  onNavigateToPage?: (page: any) => void;
  onGenerateWorkOrder?: (asset: string, desc: string) => void;
}> = ({ tele = [], onNavigateToPage, onGenerateWorkOrder }) => {
  const safeTele = (tele && tele.length > 0) ? tele : [{ vib: 11.8, temp: 94.1, psi: 242, kw: 330, flow: 84, t: '09:47' }];
  const l = safeTele[safeTele.length - 1];
  const vibTrend = safeTele.map(p => p?.vib || 0);
  const tempTrend = safeTele.map(p => p?.temp || 0);

  return (
    <div className="h-full flex flex-col bg-[#09090b]">
      <Toolbar>
        <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">OPERATIONS INTELLIGENCE</span>
        <ToolSep />
        <span className="text-[10px] text-zinc-500">EXECUTIVE SAFETY ANALYTICS & KPI BREAKDOWN</span>
        <div className="flex-1" />
        <ToolBtn><Download size={12} /> Export BI PDF</ToolBtn>
        <ToolBtn><Calendar size={12} /> Last 30 Days</ToolBtn>
      </Toolbar>

      <div className="flex-1 overflow-y-auto p-5 space-y-6">
        <div className="flex items-center justify-between p-4 rounded-xl bg-white/[0.02] border border-white/[0.04]">
          <div>
            <div className="flex items-center gap-2">
              <h2 className="text-base font-bold text-white">Executive Operations Intelligence Brief</h2>
              <span className="text-[10px] px-2 py-0.5 rounded-full bg-emerald-500/10 text-emerald-400 border border-emerald-500/20 font-semibold">
                ISO 45001 Verified
              </span>
            </div>
            <p className="text-xs text-zinc-400 mt-1">
              Plant Alpha Refinery • Shift B (143 Operators Online) • Overall Plant Safety Score: <span className="text-emerald-400 font-bold">94.2/100</span> (↑ +2.1% from yesterday).
            </p>
          </div>
          <div className="text-right text-xs text-zinc-500">
            <div>Last Updated: <strong className="text-zinc-300">14:07:36</strong></div>
            <div>Data Pipeline: <strong className="text-emerald-400">Connected</strong></div>
          </div>
        </div>

        <div className="grid grid-cols-6 gap-3">
          {[
            { label: 'Safety Score Index', val: '94.2', sub: '/100 (+2.1%)', color: 'text-emerald-400' },
            { label: 'OSHA Incident Rate', val: '0.18', sub: 'TRIR Benchmark', color: 'text-emerald-400' },
            { label: 'Active Plant Risks', val: '2', sub: 'High (DC-101)', color: 'text-amber-400' },
            { label: 'Inspection Audit Pass', val: '98.4%', sub: '122 Audited', color: 'text-emerald-400' },
            { label: 'Asset Fleet Health', val: '91.2%', sub: '47 Monitored', color: 'text-white' },
            { label: 'Permit LOTO Sync', val: '100%', sub: '28 Active PTW', color: 'text-emerald-400' },
          ].map(k => (
            <div key={k.label} className="p-3.5 rounded-xl bg-white/[0.02] border border-white/[0.04]">
              <p className="text-[9px] text-zinc-500 uppercase tracking-wider mb-1 font-semibold">{k.label}</p>
              <p className={`text-xl font-bold ${k.color}`}>{k.val}</p>
              <p className="text-[10px] text-zinc-600 mt-1">{k.sub}</p>
            </div>
          ))}
        </div>

        <div className="p-4 rounded-xl bg-indigo-600/10 border border-indigo-500/20 space-y-2">
          <div className="flex items-center gap-2">
            <Sparkles size={16} className="text-indigo-400" />
            <h3 className="text-sm font-bold text-white">AI Executive Briefing & Recommended Actions</h3>
          </div>
          <div className="grid md:grid-cols-2 gap-4 pt-1 text-xs leading-relaxed text-zinc-300">
            <div>
              <p className="font-semibold text-zinc-200 mb-1">Key Changes Today:</p>
              <ul className="space-y-1 text-zinc-400 list-disc list-inside">
                <li>Pump P-102 vibration velocity increased 3% to 11.8 mm/s.</li>
                <li>Zone B contractor audit completed; expired badge C-4412 revoked.</li>
                <li>Zero permit compliance violations across Tank T-204 isolation.</li>
                <li>AI model predicts zero production interruption over next 72 hours.</li>
              </ul>
            </div>
            <div>
              <p className="font-semibold text-zinc-200 mb-1">Executive Recommendations:</p>
              <ol className="space-y-1 text-zinc-400 list-decimal list-inside">
                <li>Approve bearing race replacement for Pump P-102 within 18 days.</li>
                <li>Review lubrication schedule auto-escalation in CMMS configuration.</li>
                <li>Maintain continuous Jetson AGX camera scan in Zone B.</li>
              </ol>
            </div>
          </div>
        </div>

        <div className="grid grid-cols-2 gap-5">
          <div className="p-4 rounded-xl bg-white/[0.015] border border-white/[0.04]">
            <div className="flex justify-between items-baseline mb-2">
              <span className="text-sm font-bold text-white">30-Day Vibration Velocity Trend — Pump P-102</span>
              <span className="text-xs text-amber-400 font-semibold">{l.vib} mm/s (Limit: 14.0)</span>
            </div>
            <Chart data={vibTrend} color="#818cf8" threshold={14} h={160} />
            <div className="flex justify-between text-[10px] text-zinc-600 mt-2">
              <span>30 Days Ago</span><span>15 Days Ago</span><span>Today (Live)</span>
            </div>
          </div>

          <div className="p-4 rounded-xl bg-white/[0.015] border border-white/[0.04]">
            <div className="flex justify-between items-baseline mb-2">
              <span className="text-sm font-bold text-white">30-Day Bearing Temperature Trend — TS-102</span>
              <span className="text-xs text-white font-semibold">{l.temp}°C (Alarm: 100°C)</span>
            </div>
            <Chart data={tempTrend} color="#f59e0b" threshold={100} h={160} />
            <div className="flex justify-between text-[10px] text-zinc-600 mt-2">
              <span>30 Days Ago</span><span>15 Days Ago</span><span>Today (Live)</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 3: EXECUTIVE INSIGHTS
// ═══════════════════════════════════════════════════════════
export const InspectionsWorkspace: React.FC = () => {
  const [filterTab, setFilterTab] = useState<'all' | 'pending' | 'overdue' | 'failed'>('all');

  const inspectionsList = [
    { id: 'AUD-9901', name: 'PPE Hardhat & Goggles Audit', area: 'Zone B North', inspector: 'Harish M.', status: 'Pending', due: 'Today (16:00)', sev: 'Medium', score: '98%' },
    { id: 'AUD-9902', name: 'Fire Safety & Hydrant Pressure', area: 'Boiler Area A', inspector: 'Priya S.', status: 'Passed', due: 'Completed', sev: 'Normal', score: '95%' },
    { id: 'AUD-9903', name: 'Electrical MCC Panel Inspection', area: 'MCC Room 7B', inspector: 'Rahul K.', status: 'Overdue', due: '2 Days Ago', sev: 'Critical', score: '88%' },
    { id: 'AUD-9904', name: 'Emergency Exit Door Clearance', area: 'Warehouse Zone C', inspector: 'John D.', status: 'Failed', due: 'Today', sev: 'High', score: '72%' },
    { id: 'AUD-9905', name: 'Hazardous Chemical Storage Audit', area: 'Tank Farm T-204', inspector: 'Suresh P.', status: 'Passed', due: 'Completed', sev: 'Normal', score: '99%' },
    { id: 'AUD-9906', name: 'Scaffolding & Rigging Safety Audit', area: 'Reactor Complex B', inspector: 'Harish M.', status: 'Pending', due: 'Tomorrow', sev: 'Medium', score: '96%' },
  ];

  const filtered = inspectionsList.filter(i => {
    if (filterTab === 'pending') return i.status === 'Pending';
    if (filterTab === 'overdue') return i.status === 'Overdue';
    if (filterTab === 'failed') return i.status === 'Failed';
    return true;
  });

  return (
    <div className="h-full flex flex-col bg-[#09090b]">
      <Toolbar>
        <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">INSPECTIONS & COMPLIANCE AUDIT CENTER</span>
        <ToolSep />
        <ToolBtn active={filterTab === 'all'} onClick={() => setFilterTab('all')}>All (139)</ToolBtn>
        <ToolBtn active={filterTab === 'pending'} onClick={() => setFilterTab('pending')}>Pending (14)</ToolBtn>
        <ToolBtn active={filterTab === 'overdue'} onClick={() => setFilterTab('overdue')}>Overdue (3)</ToolBtn>
        <ToolBtn active={filterTab === 'failed'} onClick={() => setFilterTab('failed')}>Failed (2)</ToolBtn>
        <div className="flex-1" />
        <ToolBtn className="!bg-indigo-600 !text-white"><Plus size={12} /> Schedule New Audit</ToolBtn>
        <ToolBtn><Download size={12} /> Export CSV</ToolBtn>
      </Toolbar>

      <div className="flex-1 overflow-y-auto p-5 space-y-5">
        {/* 1. Executive Summary KPIs */}
        <div className="grid grid-cols-6 gap-3">
          {[
            { l: 'Completed Today', v: '18', c: 'text-white', sub: 'On track' },
            { l: 'Completed This Month', v: '122', c: 'text-emerald-400', sub: '98.4% pass' },
            { l: 'Pending Audits', v: '14', c: 'text-amber-400', sub: 'Due in 24h' },
            { l: 'Overdue Audits', v: '3', c: 'text-red-400', sub: 'Requires escalation' },
            { l: 'Failed Checklists', v: '2', c: 'text-red-400', sub: 'CAPA generated' },
            { l: 'Compliance Score', v: '97.8%', c: 'text-emerald-400', sub: '↑ +1.2% this week' },
          ].map(k => (
            <div key={k.l} className="p-3.5 rounded-xl bg-white/[0.02] border border-white/[0.04]">
              <p className="text-[9px] text-zinc-500 uppercase tracking-wider mb-1 font-semibold">{k.l}</p>
              <p className={`text-xl font-bold ${k.c}`}>{k.v}</p>
              <p className="text-[10px] text-zinc-600 mt-1">{k.sub}</p>
            </div>
          ))}
        </div>

        {/* 2. AI Inspector Summary & Advice */}
        <div className="p-4 rounded-xl bg-indigo-600/10 border border-indigo-500/20 space-y-2">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-2">
              <Sparkles size={16} className="text-indigo-400" />
              <h3 className="text-sm font-bold text-white">AI Inspector Briefing & Recommended Action</h3>
            </div>
            <span className="text-[10px] text-zinc-400">Model: PRAHARI Inspection-v2</span>
          </div>
          <div className="grid md:grid-cols-2 gap-4 text-xs leading-relaxed text-zinc-300">
            <div>
              <p className="font-semibold text-zinc-200 mb-1">Today's Inspection Summary:</p>
              <ul className="space-y-1 text-zinc-400 list-disc list-inside">
                <li>18 audits completed today with 97.8% overall compliance score.</li>
                <li>Emergency exit blocked in Warehouse Zone C (Severity: High).</li>
                <li>2 overdue inspections in MCC Room 7B (Electrical Panel 7B).</li>
                <li>PPE compliance verified at 100% across Zone A main line.</li>
              </ul>
            </div>
            <div>
              <p className="font-semibold text-zinc-200 mb-1">Actionable Recommendation:</p>
              <p className="text-zinc-300 bg-white/[0.03] p-2.5 rounded-lg border border-white/[0.06]">
                <strong>Prioritize Warehouse Zone C emergency exit clearance</strong> before the shift change at 16:00. Dispatched EHS Lead John D. to re-inspect.
              </p>
            </div>
          </div>
        </div>

        {/* 3. Main Inspection Queue Table */}
        <div className="p-4 rounded-xl bg-white/[0.015] border border-white/[0.04] space-y-3">
          <div className="flex justify-between items-center">
            <span className="text-sm font-bold text-white">Active Audit Inspection Queue ({filtered.length})</span>
            <span className="text-xs text-zinc-500">Showing filtered results</span>
          </div>

          <div className="space-y-1.5">
            <div className="h-8 flex items-center px-4 gap-2 bg-white/[0.01] border-b border-white/[0.04] text-[10px] text-zinc-500 uppercase tracking-wider font-semibold">
              <span className="w-24">Audit ID</span>
              <span className="flex-1">Inspection Checklist Name</span>
              <span className="w-36">Plant Area / Zone</span>
              <span className="w-28">Inspector</span>
              <span className="w-24">Status</span>
              <span className="w-28">Due Window</span>
              <span className="w-16 text-right">Score</span>
            </div>

            {filtered.map(item => (
              <div key={item.id} className="h-11 flex items-center px-4 gap-2 text-[12px] hover:bg-white/[0.02] cursor-pointer transition-colors border-b border-white/[0.02]">
                <span className="w-24 text-indigo-400 font-semibold font-mono text-[11px]">{item.id}</span>
                <span className="flex-1 text-zinc-200 font-medium">{item.name}</span>
                <span className="w-36 text-zinc-400">{item.area}</span>
                <span className="w-28 text-zinc-400">{item.inspector}</span>
                <span className="w-24">
                  <span className={`text-[10px] px-2 py-0.5 rounded-full font-bold ${
                    item.status === 'Passed' ? 'bg-emerald-500/10 text-emerald-400' :
                    item.status === 'Pending' ? 'bg-amber-500/15 text-amber-400' :
                    item.status === 'Overdue' ? 'bg-red-500/15 text-red-400' : 'bg-red-500/20 text-red-300'
                  }`}>
                    {item.status}
                  </span>
                </span>
                <span className="w-28 text-zinc-500 text-[11px]">{item.due}</span>
                <span className="w-16 text-right font-bold text-white">{item.score}</span>
              </div>
            ))}
          </div>
        </div>

        {/* 4. Compliance Breakdown & Recent Findings Stream Grid */}
        <div className="grid grid-cols-2 gap-5">
          {/* Category Compliance Breakdown */}
          <div className="p-4 rounded-xl bg-white/[0.015] border border-white/[0.04] space-y-3">
            <span className="text-sm font-bold text-white block">Compliance Score by Safety Category</span>
            {[
              { cat: 'PPE & Personal Safety', score: 98, color: 'bg-emerald-500' },
              { cat: 'Fire Safety & Extinguishers', score: 95, color: 'bg-emerald-500' },
              { cat: 'Electrical MCC Systems', score: 96, color: 'bg-emerald-500' },
              { cat: 'Mechanical Rigging & Hoisting', score: 100, color: 'bg-emerald-500' },
              { cat: 'Hazardous Chemical Storage', score: 93, color: 'bg-amber-500' },
              { cat: 'Emergency Evacuation Preparedness', score: 97, color: 'bg-emerald-500' },
            ].map(c => (
              <div key={c.cat} className="space-y-1">
                <div className="flex justify-between text-xs">
                  <span className="text-zinc-300">{c.cat}</span>
                  <span className="text-zinc-400 font-bold">{c.score}%</span>
                </div>
                <div className="w-full h-1.5 rounded-full bg-zinc-800">
                  <div className={`h-1.5 rounded-full ${c.color}`} style={{ width: `${c.score}%` }} />
                </div>
              </div>
            ))}
          </div>

          {/* Live Recent Findings Log */}
          <div className="p-4 rounded-xl bg-white/[0.015] border border-white/[0.04] space-y-3">
            <span className="text-sm font-bold text-white block">Recent Audit Findings & Non-Conformances</span>
            <div className="space-y-2">
              {[
                { time: '14:10', msg: 'Emergency exit door blocked in Warehouse Zone C', sev: 'High', color: 'text-red-400' },
                { time: '13:48', msg: 'Fire extinguisher pressure low on Hydrant HY-04', sev: 'Medium', color: 'text-amber-400' },
                { time: '13:15', msg: 'PPE compliance 100% verified across Zone A Main Line', sev: 'Info', color: 'text-emerald-400' },
                { time: '12:30', msg: 'Cable insulation wear detected on MCC-7B Panel', sev: 'High', color: 'text-red-400' },
              ].map((f, i) => (
                <div key={i} className="p-2.5 rounded-lg bg-white/[0.02] border border-white/[0.04] text-xs">
                  <div className="flex justify-between items-center mb-1">
                    <span className="font-mono text-zinc-500 text-[10px]">{f.time}</span>
                    <span className={`text-[10px] font-bold ${f.color}`}>{f.sev} Severity</span>
                  </div>
                  <span className="text-zinc-300 font-medium">{f.msg}</span>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export const ExecutiveInsights: React.FC = () => (
  <div className="h-full flex flex-col bg-[#09090b]">
    <Toolbar>
      <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">EXECUTIVE COMPLIANCE INSIGHTS</span>
      <ToolSep />
      <span className="text-[10px] text-zinc-500">ISO 45001 & OSHA COMPLIANCE AUDIT CENTER</span>
      <div className="flex-1" />
      <ToolBtn className="!bg-indigo-600 !text-white"><Download size={12} /> Export Full Audit Package</ToolBtn>
    </Toolbar>

    <div className="flex-1 overflow-y-auto p-5 space-y-6">
      <div className="flex items-center justify-between p-4 rounded-xl bg-white/[0.02] border border-white/[0.04]">
        <div>
          <div className="flex items-center gap-2">
            <h2 className="text-base font-bold text-white">Executive Compliance Score: <span className="text-emerald-400">97.6%</span></h2>
            <span className="text-[10px] px-2 py-0.5 rounded-full bg-emerald-500/10 text-emerald-400 border border-emerald-500/20 font-semibold">
              Status: GOOD (↑ +1.4% this month)
            </span>
          </div>
          <p className="text-xs text-zinc-400 mt-1">
            Last Audit: <strong className="text-zinc-200">3 Days Ago</strong> • 1,284 Evidence Files Verified • Zero Open Regulatory Enforcement Notices.
          </p>
        </div>
        <div className="text-right text-xs text-zinc-500">
          <div>Verified By: <strong className="text-indigo-400">Compliance Agent</strong></div>
          <div>Audit Hash: <strong className="text-zinc-400 font-mono">0x4a91...b82</strong></div>
        </div>
      </div>

      <div className="grid grid-cols-5 gap-3">
        {[
          { std: 'ISO 45001', score: '98%', clauses: '42/45 Compliant', color: 'text-emerald-400' },
          { std: 'OSHA 29 CFR', score: '95%', clauses: '118/120 Req Met', color: 'text-emerald-400' },
          { std: 'Internal EHS SOP', score: '100%', clauses: 'Completed', color: 'text-emerald-400' },
          { std: 'Environmental', score: '96%', clauses: 'EPA Standard', color: 'text-emerald-400' },
          { std: 'Contractor Audit', score: '91%', clauses: '1 Expired Badge', color: 'text-amber-400' },
        ].map(s => (
          <div key={s.std} className="p-3.5 rounded-xl bg-white/[0.02] border border-white/[0.04]">
            <p className="text-[10px] font-semibold text-zinc-400 uppercase tracking-wider mb-1">{s.std}</p>
            <p className={`text-xl font-bold ${s.color}`}>{s.score}</p>
            <p className="text-[10px] text-zinc-500 mt-1">{s.clauses}</p>
          </div>
        ))}
      </div>

      <div className="p-4 rounded-xl bg-indigo-600/10 border border-indigo-500/20 space-y-2">
        <div className="flex items-center gap-2">
          <Sparkles size={16} className="text-indigo-400" />
          <h3 className="text-sm font-bold text-white">AI Compliance Summary & Action Plan</h3>
        </div>
        <div className="grid md:grid-cols-2 gap-4 pt-1 text-xs leading-relaxed text-zinc-300">
          <div>
            <p className="font-semibold text-zinc-200 mb-1">Today's Verified Changes:</p>
            <ul className="space-y-1 text-zinc-400 list-disc list-inside">
              <li>2 contractor certifications expire this week (Badge C-4412 auto-revoked).</li>
              <li>Emergency drill due in 7 days (Warehouse Zone C).</li>
              <li>PPE compliance index increased from 94% to 97% across Zone B.</li>
              <li>ISO Clause 8.1 operational planning documentation verified.</li>
            </ul>
          </div>
          <div>
            <p className="font-semibold text-zinc-200 mb-1">Recommended Corrective Actions:</p>
            <ol className="space-y-1 text-zinc-400 list-decimal list-inside">
              <li>Schedule Warehouse emergency response drill before Friday.</li>
              <li>Renew contractor medical certification for C-4412.</li>
              <li>Upload missing permit attachment for PTW-8903.</li>
            </ol>
          </div>
        </div>
      </div>

      <div className="p-4 rounded-xl bg-white/[0.015] border border-white/[0.04] space-y-3">
        <span className="text-sm font-bold text-white block">Plant Area Compliance Score Heatmap</span>
        <div className="grid grid-cols-5 gap-3">
          {[
            { zone: 'Zone A (Main Line)', score: '98%', status: 'Compliant', color: 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20' },
            { zone: 'Zone B (Reactor North)', score: '92%', status: 'Attention', color: 'bg-amber-500/10 text-amber-400 border-amber-500/20' },
            { zone: 'Tank Farm T-204', score: '100%', status: 'Optimal', color: 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20' },
            { zone: 'Utilities & Steam', score: '95%', status: 'Compliant', color: 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20' },
            { zone: 'Warehouse Storage', score: '89%', status: 'Drill Due', color: 'bg-amber-500/10 text-amber-400 border-amber-500/20' },
          ].map(z => (
            <div key={z.zone} className={`p-3 rounded-xl border ${z.color} text-xs space-y-1`}>
              <span className="font-bold text-white block">{z.zone}</span>
              <span className="text-lg font-extrabold">{z.score}</span>
              <span className="block text-[10px] opacity-80">{z.status}</span>
            </div>
          ))}
        </div>
      </div>
    </div>
  </div>
);

// ═══════════════════════════════════════════════════════════
// WORKSPACE 4: OPERATIONS CENTER
// ═══════════════════════════════════════════════════════════
export const OperationsCenter: React.FC<{ tele?: TelemetryPoint[]; onNavigateToAsset?: (assetId: string) => void }> = ({ tele = [], onNavigateToAsset }) => {
  const safeTele = (tele && tele.length > 0) ? tele : [{ vib: 11.8, temp: 94.1, psi: 242, kw: 330, flow: 84, t: '09:47' }];
  const l = safeTele[safeTele.length - 1];
  const [chartMode, setChartMode] = useState<'vib' | 'temp' | 'psi' | 'kw'>('vib');
  const chartConfig = {
    vib: { data: safeTele.map(p => p?.vib || 0), color: '#818cf8', threshold: 14, label: 'Vibration Velocity', unit: 'mm/s', alarm: 'ISO 10816 Limit @ 14.0' },
    temp: { data: safeTele.map(p => p?.temp || 0), color: '#f59e0b', threshold: 100, label: 'Bearing Temperature', unit: '°C', alarm: 'Alarm @ 100°C' },
    psi: { data: safeTele.map(p => p?.psi || 0), color: '#22d3ee', threshold: 260, label: 'Line Pressure', unit: 'PSI', alarm: 'High @ 260 PSI' },
    kw: { data: safeTele.map(p => +(p?.kw || 0)), color: '#a78bfa', threshold: 370, label: 'Power Draw', unit: 'kW', alarm: 'Peak @ 370 kW' },
  }[chartMode];

  return (
    <div className="h-full flex flex-col bg-[#09090b]">
      <Toolbar>
        <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">OPERATIONS CENTER</span>
        <ToolSep />
        <ToolBtn active={chartMode === 'vib'} onClick={() => setChartMode('vib')}>Vibration</ToolBtn>
        <ToolBtn active={chartMode === 'temp'} onClick={() => setChartMode('temp')}>Temperature</ToolBtn>
        <ToolBtn active={chartMode === 'psi'} onClick={() => setChartMode('psi')}>Pressure</ToolBtn>
        <ToolBtn active={chartMode === 'kw'} onClick={() => setChartMode('kw')}>Power</ToolBtn>
        <div className="flex-1" />
        <span className="text-[10px] text-emerald-400 font-medium flex items-center gap-1"><span className="w-1.5 h-1.5 rounded-full bg-emerald-400 animate-pulse" /> Live Telemetry Stream</span>
      </Toolbar>

      <div className="flex-1 flex overflow-hidden">
        <div className="flex-1 flex flex-col overflow-hidden">
          <div className="h-16 flex items-center gap-8 px-5 border-b border-white/[0.04] bg-white/[0.01]">
            <Metric label="Vibration" value={l.vib} unit="mm/s" accent={l.vib > 12 ? 'text-amber-400' : undefined} small />
            <Metric label="Temperature" value={l.temp} unit="°C" accent={l.temp > 95 ? 'text-amber-400' : undefined} small />
            <Metric label="Pressure" value={l.psi} unit="PSI" small />
            <Metric label="Power" value={l.kw} unit="kW" small />
            <Metric label="Flow" value={l.flow} unit="L/m" small />
          </div>

          <div className="flex-1 p-4 flex flex-col">
            <div className="flex items-baseline justify-between mb-2">
              <div>
                <span className="text-sm font-medium text-white">{chartConfig.label} — Pump P-102</span>
                <span className="text-xs text-zinc-600 ml-2">{chartConfig.alarm}</span>
              </div>
              <span className="text-lg font-semibold text-white">{chartConfig.data[chartConfig.data.length - 1]} <span className="text-xs text-zinc-500 font-normal">{chartConfig.unit}</span></span>
            </div>
            <div className="flex-1 rounded-xl bg-white/[0.015] border border-white/[0.04] p-3">
              <Chart data={chartConfig.data} color={chartConfig.color} threshold={chartConfig.threshold} h={220} />
            </div>
          </div>
        </div>

        <div className="w-72 border-l border-white/[0.04] bg-white/[0.01] flex flex-col shrink-0">
          <div className="h-9 px-4 border-b border-white/[0.04] flex items-center justify-between">
            <span className="text-[10px] font-semibold text-zinc-400 uppercase tracking-wider">Equipment Status</span>
            <span className="text-[10px] text-zinc-600">47 Total</span>
          </div>
          <div className="flex-1 overflow-y-auto p-3 space-y-2">
            {[
              { name: 'Pump P-102', health: 74, st: 'Warning', temp: l.temp, vib: l.vib },
              { name: 'Valve V-88', health: 98, st: 'Locked', temp: '--', vib: '--' },
              { name: 'HX-04', health: 91, st: 'Running', temp: '88', vib: '3.2' },
              { name: 'Compressor C-03', health: 96, st: 'Running', temp: '72', vib: '4.1' },
              { name: 'Boiler A', health: 93, st: 'Running', temp: '210', vib: '2.8' },
              { name: 'Reactor B', health: 97, st: 'Running', temp: '340', vib: '1.9' },
            ].map(eq => (
              <div key={eq.name} className="p-3 rounded-xl bg-white/[0.02] hover:bg-white/[0.04] border border-white/[0.04] cursor-pointer transition-colors">
                <div className="flex items-center justify-between mb-1">
                  <span className="text-xs font-semibold text-white">{eq.name}</span>
                  <span className={`w-2 h-2 rounded-full ${eq.health > 90 ? 'bg-emerald-400' : eq.health > 80 ? 'bg-amber-400' : 'bg-red-400'}`} />
                </div>
                <div className="flex justify-between text-[10px] text-zinc-500">
                  <span>Health: {eq.health}%</span>
                  <span className={eq.st === 'Warning' ? 'text-amber-400 font-semibold' : 'text-zinc-400'}>{eq.st}</span>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 5: INDUSTRIAL TWIN
// ═══════════════════════════════════════════════════════════
export const IndustrialTwin: React.FC<{ tele?: TelemetryPoint[]; initialSelectedAssetId?: string }> = ({ tele = [], initialSelectedAssetId }) => {
  const safeTele = (tele && tele.length > 0) ? tele : [{ temp: 94.1, vib: 11.8 }];
  const l = safeTele[safeTele.length - 1];
  const [selected, setSelected] = useState(initialSelectedAssetId || 'pump');
  const assets = [
    { id: 'reactor', label: 'Reactor B', x: 180, y: 130, health: 97, temp: '340°C', st: 'Running', rul: '94d' },
    { id: 'pump', label: 'Pump P-102', x: 400, y: 200, health: 74, temp: `${l.temp}°C`, st: 'Warning', rul: '18d' },
    { id: 'hx', label: 'HX-04', x: 600, y: 120, health: 91, temp: '88°C', st: 'Running', rul: '62d' },
    { id: 'comp', label: 'Compressor C-03', x: 360, y: 340, health: 96, temp: '72°C', st: 'Running', rul: '84d' },
    { id: 'valve', label: 'Valve V-88', x: 140, y: 310, health: 98, temp: 'N/A', st: 'Locked', rul: 'N/A' },
    { id: 'boiler', label: 'Boiler A', x: 620, y: 310, health: 93, temp: '210°C', st: 'Running', rul: '71d' },
  ];
  const sel = assets.find(a => a.id === selected);

  return (
    <div className="h-full flex flex-col bg-[#09090b]">
      <Toolbar>
        <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">INDUSTRIAL TWIN</span>
        <ToolSep />
        <span className="text-[10px] text-zinc-500">Reactor Complex B CAD/SCADA Interactive Schematic</span>
      </Toolbar>

      <div className="flex-1 flex overflow-hidden">
        <div className="flex-1 relative bg-[#0b0b0f]">
          <svg viewBox="0 0 800 460" className="w-full h-full" preserveAspectRatio="xMidYMid meet">
            {[[240,150,380,200],[460,200,580,140],[400,240,380,320],[180,190,160,290],[430,340,580,320]].map(([x1,y1,x2,y2], i) => (
              <line key={i} x1={x1} y1={y1} x2={x2} y2={y2} stroke="rgba(99,102,241,0.15)" strokeWidth="2" strokeDasharray="8,6" />
            ))}
            {assets.map(a => {
              const s = selected === a.id;
              const w = a.st === 'Warning';
              return (
                <g key={a.id} onClick={() => setSelected(a.id)} className="cursor-pointer">
                  <circle cx={a.x} cy={a.y} r={s ? 40 : 34} fill={s ? 'rgba(99,102,241,0.15)' : w ? 'rgba(245,158,11,0.08)' : 'rgba(255,255,255,0.02)'} stroke={s ? '#6366f1' : w ? '#f59e0b' : 'rgba(255,255,255,0.08)'} strokeWidth={s ? 2 : 1} />
                  <circle cx={a.x} cy={a.y} r={5} fill={a.health > 90 ? '#22c55e' : a.health > 80 ? '#f59e0b' : '#ef4444'} />
                  <text x={a.x} y={a.y + 54} textAnchor="middle" fill={s ? '#c7d2fe' : 'rgba(255,255,255,0.4)'} fontSize="11" fontFamily="Inter" fontWeight="500">{a.label}</text>
                </g>
              );
            })}
          </svg>
        </div>

        {sel && (
          <div className="w-80 border-l border-white/[0.04] p-4 bg-white/[0.01] space-y-4 overflow-y-auto">
            <div>
              <p className="text-[10px] text-zinc-500 uppercase">Selected Equipment</p>
              <h2 className="text-base font-bold text-white mt-0.5">{sel.label}</h2>
              <span className={`inline-block mt-1 text-[10px] px-2 py-0.5 rounded-full ${sel.st === 'Warning' ? 'bg-amber-500/15 text-amber-400 font-semibold' : 'bg-emerald-500/10 text-emerald-400'}`}>{sel.st}</span>
            </div>
            <div className="grid grid-cols-2 gap-3">
              <Metric label="Health Score" value={`${sel.health}%`} small />
              <Metric label="RUL Forecast" value={sel.rul} accent={sel.rul === '18d' ? 'text-amber-400' : undefined} small />
              <Metric label="Temperature" value={sel.temp} small />
              <Metric label="Status" value={sel.st} small />
            </div>

            {selected === 'pump' && (
              <div className="p-3.5 rounded-xl bg-indigo-600/10 border border-indigo-500/20 space-y-2">
                <div className="flex items-center gap-1.5">
                  <Sparkles size={14} className="text-indigo-400" />
                  <span className="text-[10px] font-bold text-indigo-300 uppercase tracking-wider">PINN Neural Model Failure Analysis</span>
                </div>
                <p className="text-[11px] text-zinc-300 leading-relaxed">
                  Bearing outer race wear probability: <strong>72%</strong>. RUL estimate: 18 days. Lubrication service interval was exceeded by 14 days under thermal load.
                </p>
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 6: AI COMMAND CENTER
// ═══════════════════════════════════════════════════════════
export const AICommandCenter: React.FC<{ onGenerateWorkOrder?: (asset: string, desc: string) => void }> = ({ onGenerateWorkOrder }) => {
  const steps = [
    { text: 'Detected vibration anomaly on Pump P-102', status: 'done', detail: 'Reading 11.8 mm/s, trending toward ISO limit.' },
    { text: 'Queried CMMS maintenance history from PostgreSQL', status: 'done', detail: 'PM overdue by 14 days. Work order WO-7821 un-escalated.' },
    { text: 'Calculated bearing wear probability in Digital Twin', status: 'done', detail: 'PINN model predicts 72% wear probability. RUL: 18 days.' },
    { text: 'Verified PTW-8902 LOTO isolation lock on Valve V-88', status: 'done', detail: 'Physical gate valve lock confirmed. Zero safety conflicts.' },
    { text: 'Synthesizing corrective action plan (CAPA)', status: 'active', detail: '4 actions generated for maintenance crew.' },
  ];

  return (
    <div className="h-full flex flex-col bg-[#09090b]">
      <Toolbar>
        <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">AI COMMAND CENTER</span>
        <ToolSep />
        <span className="text-[10px] text-indigo-400 font-medium">10 Autonomous Agents Active</span>
        <div className="flex-1" />
        <ToolBtn onClick={() => onGenerateWorkOrder?.('PUMP-P102', 'Bearing race replacement')} className="!bg-indigo-600 !text-white">
          <Wrench size={12} /> Dispatch WO-7821
        </ToolBtn>
      </Toolbar>

      <div className="flex-1 flex overflow-hidden">
        <div className="flex-1 p-5 space-y-4 overflow-y-auto">
          <div className="p-4 rounded-xl bg-indigo-600/10 border border-indigo-500/20 flex justify-between items-center">
            <div>
              <h2 className="text-sm font-bold text-white">Active Reasoning Stream: Pump P-102 Anomaly</h2>
              <p className="text-xs text-zinc-400 mt-1">Multi-agent supervisor is conducting 5-Whys RCA, evidence correlation, and ISO 45001 compliance audit.</p>
            </div>
            <button
              onClick={() => onGenerateWorkOrder?.('PUMP-P102', 'Bearing race replacement')}
              className="px-3 py-1.5 rounded-lg bg-indigo-600 hover:bg-indigo-500 text-white font-bold text-xs shadow-lg shadow-indigo-600/30"
            >
              Dispatch Work Order WO-7821
            </button>
          </div>
          <div className="space-y-2">
            {steps.map((s, i) => (
              <div key={i} className="flex gap-3 py-3 px-3 rounded-xl bg-white/[0.02] border border-white/[0.04]">
                <CheckCircle2 size={16} className={s.status === 'done' ? 'text-emerald-400 mt-0.5' : 'text-indigo-400 mt-0.5'} />
                <div>
                  <p className="text-[13px] font-semibold text-zinc-200">{s.text}</p>
                  <p className="text-[12px] text-zinc-500 mt-0.5">{s.detail}</p>
                </div>
              </div>
            ))}
          </div>
        </div>

        <div className="w-80 border-l border-white/[0.04] p-4 bg-white/[0.01] space-y-4">
          <span className="text-[10px] font-semibold text-zinc-400 uppercase tracking-wider block">Agent DAG Status</span>
          {[
            { n: 'Supervisor', r: 'Orchestrator', s: 'active' },
            { n: 'Inspection', r: 'Vision Audit', s: 'done' },
            { n: 'Risk Assessment', r: '5×5 Matrix', s: 'done' },
            { n: 'Permit', r: 'PTW Isolation', s: 'done' },
            { n: 'Maintenance', r: 'Digital Twin', s: 'active' },
            { n: 'Incident', r: 'Root Cause', s: 'done' },
          ].map(a => (
            <div key={a.n} className="flex items-center justify-between text-xs">
              <span className="text-zinc-300">{a.n}</span>
              <span className={`text-[10px] font-semibold ${a.s === 'active' ? 'text-indigo-400' : 'text-emerald-400'}`}>{a.s === 'active' ? '● Running' : '✓ Done'}</span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 7: INCIDENTS WORKSPACE
// ═══════════════════════════════════════════════════════════
export const IncidentsWorkspace: React.FC<{ incidents?: Incident[]; onReportIncident?: () => void }> = ({ incidents = [], onReportIncident }) => {
  const defaultIncidents: Incident[] = [
    { id: 'INC-2026-0447', title: 'Pump P-102 Vibration Anomaly', desc: 'Vibration probe threshold exceeded during high-pressure run.', sev: 'Warning', asset: 'Pump P-102', st: 'Under Investigation', time: '14:10' },
    { id: 'INC-2026-0442', title: 'Contractor Badge Expiration Zone B', desc: 'Contractor C-4412 detected with expired safety badge.', sev: 'Info', asset: 'Gate B', st: 'Resolved', time: '12:30' },
  ];
  const displayIncidents = incidents.length > 0 ? incidents : defaultIncidents;

  return (
    <div className="h-full flex flex-col bg-[#09090b]">
      <Toolbar>
        <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">INCIDENTS & HAZARD INVESTIGATION</span>
        <div className="flex-1" />
        <ToolBtn onClick={onReportIncident} className="!bg-red-600/80 !text-white">
          <Plus size={12} /> Report Incident
        </ToolBtn>
      </Toolbar>
      <div className="flex-1 flex overflow-hidden">
        <div className="flex-1 p-5 space-y-4 overflow-y-auto">
          {displayIncidents.map(inc => (
            <div key={inc.id} className="p-4 rounded-xl bg-white/[0.02] border border-white/[0.04]">
              <div className="flex justify-between items-center mb-2">
                <span className="text-xs font-mono text-indigo-400 font-bold">{inc.id}</span>
                <span className="text-xs px-2 py-0.5 rounded-full bg-amber-500/15 text-amber-400 font-semibold">{inc.st}</span>
              </div>
              <h2 className="text-base font-bold text-white">{inc.title || inc.desc}</h2>
              <p className="text-xs text-zinc-400 mt-1">Asset: <span className="text-zinc-200 font-semibold">{inc.asset}</span> • Severity: <span className="text-amber-400 font-bold">{inc.sev}</span> • Logged: {inc.time || 'Today'}</p>
            </div>
          ))}
        </div>

      <div className="w-80 border-l border-white/[0.04] p-4 bg-white/[0.01] space-y-3">
        <span className="text-[10px] font-semibold text-zinc-400 uppercase tracking-wider block">Corrective Actions (CAPA)</span>
        {[
          { a: 'Replace bearing assembly P-102', o: 'Mech Team', d: '24h', s: 'In Progress' },
          { a: 'Recalibrate vibration probe VP-102', o: 'I&C Team', d: '48h', s: 'Assigned' },
          { a: 'Configure PM auto-escalation in CMMS', o: 'CMMS Admin', d: '7d', s: 'Pending' },
          { a: 'Zone B contractor badge audit', o: 'EHS Team', d: '24h', s: 'Completed' },
        ].map((c, i) => (
          <div key={i} className="p-2.5 rounded-lg bg-white/[0.02] border border-white/[0.04] text-xs">
            <span className="font-medium text-white block">{c.a}</span>
            <span className="text-[10px] text-zinc-500">{c.o} • Due: {c.d} • {c.s}</span>
          </div>
        ))}
      </div>
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 8: SAFE WORK PERMITS
// ═══════════════════════════════════════════════════════════
export const PermitsWorkspace: React.FC = () => (
  <div className="h-full flex flex-col bg-[#09090b]">
    <Toolbar>
      <span className="text-[11px] font-semibold text-zinc-300">SAFE WORK PERMITS (PTW / LOTO)</span>
      <ToolSep />
      <ToolBtn active>28 Active Permits</ToolBtn>
      <div className="flex-1" />
      <ToolBtn className="!bg-indigo-600 !text-white"><Plus size={12} /> Issue Safe Work Permit</ToolBtn>
    </Toolbar>
    <div className="flex-1 p-5 space-y-3 overflow-y-auto">
      <div className="h-8 flex items-center px-4 gap-2 bg-white/[0.01] border-b border-white/[0.04] text-[10px] text-zinc-500 uppercase tracking-wider font-semibold">
        <span className="w-24">Permit ID</span><span className="flex-1">Description</span><span className="w-24">Type</span><span className="w-28">Isolation</span><span className="w-24">Gas Test</span><span className="w-20">Status</span>
      </div>
      {[
        { id: 'PTW-8902', desc: 'Hot Work — Tank T-204 Isolation Lock', type: 'Hot Work', iso: 'V-88 LOCKED', gas: 'O₂ 20.9%', st: 'Approved' },
        { id: 'PTW-8903', desc: 'Confined Entry — Vessel V-109 Cleaning', type: 'Confined', iso: 'LINE BLIND', gas: 'Pending', st: 'Hold' },
        { id: 'PTW-8904', desc: 'Electrical Maintenance — Panel MCC-7B', type: 'Electrical', iso: 'RACKED OUT', gas: 'N/A', st: 'Approved' },
        { id: 'PTW-8905', desc: 'Excavation — Pipe Trench Line B', type: 'Excavation', iso: 'N/A', gas: 'N/A', st: 'Pending' },
      ].map(p => (
        <div key={p.id} className="h-11 flex items-center px-4 gap-2 text-[12px] hover:bg-white/[0.02] cursor-pointer transition-colors border-b border-white/[0.02]">
          <span className="w-24 text-indigo-400 font-semibold font-mono text-[11px]">{p.id}</span>
          <span className="flex-1 text-zinc-200 font-medium">{p.desc}</span>
          <span className="w-24 text-zinc-400">{p.type}</span>
          <span className="w-28 text-zinc-400 font-mono text-[11px]">{p.iso}</span>
          <span className="w-24 text-zinc-400">{p.gas}</span>
          <span className="w-20"><span className={`text-[10px] px-2 py-0.5 rounded-full font-semibold ${p.st === 'Approved' ? 'bg-emerald-500/10 text-emerald-400' : 'bg-amber-500/15 text-amber-400'}`}>{p.st}</span></span>
        </div>
      ))}
    </div>
  </div>
);

// ═══════════════════════════════════════════════════════════
// WORKSPACE 9: MAINTENANCE WORKFLOW
// ═══════════════════════════════════════════════════════════
export const MaintenanceWorkspace: React.FC<{ workOrders?: WorkOrder[] }> = ({ workOrders = [] }) => {
  const defaultWos: WorkOrder[] = [
    { id: 'WO-7821', desc: 'Bearing replacement and lubrication service', asset: 'P-102', pri: 'Critical', rul: '18d', st: 'Overdue' },
    { id: 'WO-7822', desc: 'Vibration probe recalibration', asset: 'P-102', pri: 'High', rul: '18d', st: 'Assigned' },
    { id: 'WO-7823', desc: 'Quarterly compressor inspection', asset: 'C-03', pri: 'Medium', rul: '84d', st: 'Scheduled' },
    { id: 'WO-7824', desc: 'Boiler tube thickness measurement', asset: 'Boiler A', pri: 'Medium', rul: '71d', st: 'Scheduled' },
  ];
  const displayWos = workOrders.length > 0 ? workOrders : defaultWos;

  return (
    <div className="h-full flex flex-col bg-[#09090b]">
      <Toolbar><span className="text-[11px] font-semibold text-zinc-300">PREDICTIVE MAINTENANCE WORKFLOW</span></Toolbar>
      <div className="flex-1 p-5 space-y-3 overflow-y-auto">
        <div className="h-8 flex items-center px-4 gap-2 bg-white/[0.01] border-b border-white/[0.04] text-[10px] text-zinc-500 uppercase tracking-wider font-semibold">
          <span className="w-24">WO ID</span><span className="flex-1">Description</span><span className="w-24">Asset</span><span className="w-16">Priority</span><span className="w-16">RUL</span><span className="w-20">Status</span>
        </div>
        {displayWos.map((w, idx) => (
          <div key={w.id + idx} className="h-11 flex items-center px-4 gap-2 text-[12px] hover:bg-white/[0.02] cursor-pointer transition-colors border-b border-white/[0.02]">
            <span className="w-24 text-indigo-400 font-semibold font-mono text-[11px]">{w.id}</span>
            <span className="flex-1 text-zinc-200 font-medium">{w.desc}</span>
            <span className="w-24 text-zinc-400">{w.asset}</span>
            <span className="w-16"><span className={`text-[10px] font-bold ${w.pri === 'Critical' ? 'text-red-400' : 'text-amber-400'}`}>{w.pri}</span></span>
            <span className="w-16 text-amber-400 font-semibold">{w.rul}</span>
            <span className="w-20"><span className={`text-[10px] px-2 py-0.5 rounded-full font-semibold ${w.st === 'Overdue' ? 'bg-red-500/15 text-red-400' : 'bg-indigo-500/15 text-indigo-400'}`}>{w.st}</span></span>
          </div>
        ))}
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 10: RISK ASSESSMENT
// ═══════════════════════════════════════════════════════════
export const RiskWorkspace: React.FC = () => (
  <div className="h-full flex flex-col bg-[#09090b]">
    <Toolbar><span className="text-[11px] font-semibold text-zinc-300">INDUSTRIAL RISK MATRIX & REGISTER</span></Toolbar>
    <div className="p-5 flex gap-6">
      <div className="flex-1">
        <p className="text-[11px] text-zinc-400 font-bold uppercase tracking-wider mb-3">5×5 Risk Assessment Matrix</p>
        <div className="grid grid-cols-6 gap-px bg-white/[0.04] rounded-xl overflow-hidden">
          <div className="bg-[#09090b] p-2" />{['Rare', 'Unlikely', 'Possible', 'Likely', 'Almost Certain'].map(h => <div key={h} className="bg-[#0d0d12] p-2 text-[9px] text-zinc-500 text-center font-bold">{h}</div>)}
          {['Catastrophic', 'Major', 'Moderate', 'Minor', 'Insignificant'].map((sev, si) => (
            <React.Fragment key={sev}>
              <div className="bg-[#0d0d12] p-2 text-[9px] text-zinc-500 text-right font-bold">{sev}</div>
              {[1,2,3,4,5].map(li => {
                const risk = (5 - si) * li;
                const isHighlighted = si === 1 && li === 3;
                return (
                  <div key={li} className={`p-2 text-center text-[11px] font-bold ${risk > 15 ? 'bg-red-500/15 text-red-400' : risk > 10 ? 'bg-amber-500/15 text-amber-400' : 'bg-emerald-500/5 text-emerald-400/60'} ${isHighlighted ? 'ring-2 ring-indigo-500' : ''}`}>
                    {risk}{isHighlighted && <span className="block text-[8px] text-indigo-300">P-102</span>}
                  </div>
                );
              })}
            </React.Fragment>
          ))}
        </div>
      </div>
    </div>
  </div>
);

// ═══════════════════════════════════════════════════════════
// WORKSPACE 11: VISION INTELLIGENCE
// ═══════════════════════════════════════════════════════════
export const VisionIntelligence: React.FC = () => (
  <div className="h-full flex flex-col bg-[#09090b]">
    <Toolbar><span className="text-[11px] font-semibold text-zinc-300">VISION INTELLIGENCE (YOLOv8 EDGE)</span></Toolbar>
    <div className="p-5 grid grid-cols-2 gap-4">
      {[
        { cam: 'CAM-001', loc: 'Main Assembly Line', det: 'PPE Verified — 3 Workers', ok: true },
        { cam: 'CAM-002', loc: 'Reactor B North', det: 'NO HELMET — Restricted Zone', ok: false },
        { cam: 'CAM-003', loc: 'Boiler Room A', det: 'Thermal 62°C Nominal', ok: true },
        { cam: 'CAM-004', loc: 'Loading Bay North', det: 'Forklift Proximity Safe', ok: true },
      ].map(c => (
        <div key={c.cam} className="p-4 rounded-xl bg-white/[0.02] border border-white/[0.04]">
          <div className="flex justify-between items-center mb-2">
            <span className="text-sm font-bold text-white">{c.cam} — {c.loc}</span>
            <span className={`text-[10px] font-bold px-2 py-0.5 rounded-full ${c.ok ? 'bg-emerald-500/10 text-emerald-400' : 'bg-red-500/15 text-red-400 animate-pulse'}`}>{c.ok ? 'CLEAR' : 'VIOLATION'}</span>
          </div>
          <p className="text-xs text-zinc-400">{c.det}</p>
        </div>
      ))}
    </div>
  </div>
);

// ═══════════════════════════════════════════════════════════
// WORKSPACE 12: AGENT ORCHESTRATION
// ═══════════════════════════════════════════════════════════
export const AgentOrchestration: React.FC = () => (
  <div className="h-full flex flex-col bg-[#09090b]">
    <Toolbar><span className="text-[11px] font-semibold text-zinc-300">10-AGENT MULTI-AGENT DAG ORCHESTRATION</span></Toolbar>
    <div className="p-5 space-y-2 overflow-y-auto">
      {[
        { n: 'Supervisor Agent', r: 'Orchestrator', lat: '12ms', conf: '99%', s: 'Running' },
        { n: 'Inspection Agent', r: 'Vision Audit', lat: '24ms', conf: '96%', s: 'Done' },
        { n: 'Risk Assessment Agent', r: '5×5 Matrix', lat: '18ms', conf: '97%', s: 'Done' },
        { n: 'Permit Agent', r: 'PTW Isolation', lat: '15ms', conf: '98%', s: 'Done' },
        { n: 'Maintenance Agent', r: 'Digital Twin & RUL', lat: '32ms', conf: '94%', s: 'Running' },
      ].map((a, i) => (
        <div key={i} className="flex items-center justify-between p-3 rounded-xl bg-white/[0.02] border border-white/[0.04] text-xs">
          <div><span className="font-bold text-white">{a.n}</span><span className="text-zinc-500 ml-2">({a.r})</span></div>
          <div className="flex gap-4 text-zinc-400"><span>Latency: {a.lat}</span><span>Confidence: {a.conf}</span><span className="text-indigo-400 font-semibold">{a.s}</span></div>
        </div>
      ))}
    </div>
  </div>
);

// ═══════════════════════════════════════════════════════════
// WORKSPACE 13: ASSETS WORKSPACE
// ═══════════════════════════════════════════════════════════
export const AssetsWorkspace: React.FC<{ assets?: Asset[]; tele?: TelemetryPoint[]; onAddAsset?: () => void; onNavigateToAsset?: (assetId: string) => void }> = ({ assets = [], onAddAsset, onNavigateToAsset }) => {
  const defaultAssets = [
    { name: 'Pump P-102', loc: 'DC-101 Recirc', type: 'Centrifugal Pump', health: 74, rul: '18d', st: 'Warning', owner: 'Mechanical Team', vib: 11.8, temp: 94.1 },
    { name: 'Valve V-88', loc: 'DC-101 Isol', type: 'Gate Valve', health: 98, rul: 'N/A', st: 'Locked', owner: 'Safety Team', vib: '--', temp: '--' },
    { name: 'Heat Exchanger HX-04', loc: 'DC-101 Cool', type: 'Shell & Tube', health: 91, rul: '62d', st: 'Running', owner: 'Process Team', vib: 3.2, temp: 88 },
    { name: 'Compressor C-03', loc: 'DC-102', type: 'Reciprocating', health: 96, rul: '84d', st: 'Running', owner: 'Mechanical Team', vib: 4.1, temp: 72 },
  ];
  const displayAssets = assets.length > 0 ? assets : defaultAssets;

  return (
    <div className="h-full flex flex-col bg-[#09090b]">
      <Toolbar>
        <span className="text-[11px] font-semibold text-zinc-300">ASSET REGISTRY (IBM MAXIMO FLEET)</span>
        <div className="flex-1" />
        <ToolBtn onClick={onAddAsset} className="!bg-indigo-600 !text-white"><Plus size={12} /> Add Asset</ToolBtn>
      </Toolbar>
      <div className="p-5 space-y-3 overflow-y-auto">
        {displayAssets.map((as, idx) => (
          <div key={as.name + idx} onClick={() => onNavigateToAsset?.(as.name.toLowerCase())} className="flex items-center justify-between p-3.5 rounded-xl bg-white/[0.02] hover:bg-white/[0.04] border border-white/[0.04] text-xs cursor-pointer transition-colors">
            <div>
              <span className="font-bold text-white">{as.name}</span>
              <span className="text-zinc-500 ml-3">{as.loc}</span>
              {as.type && <span className="text-[10px] text-zinc-500 ml-2">({as.type})</span>}
            </div>
            <div className="flex gap-6">
              <span className={as.health < 80 ? 'text-amber-400 font-bold' : 'text-emerald-400'}>Health: {as.health}%</span>
              <span className="text-zinc-400">RUL: {as.rul}</span>
              <span className="text-indigo-400 font-semibold">{as.st}</span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 14: PLATFORM OPERATIONS
// ═══════════════════════════════════════════════════════════
export const PlatformOps: React.FC<{ health: HealthStatus }> = ({ health }) => (
  <div className="h-full flex flex-col bg-[#09090b]">
    <Toolbar><span className="text-[11px] font-semibold text-zinc-300">PLATFORM OPERATIONS (AWS OBSERVABILITY)</span></Toolbar>
    <div className="flex-1 overflow-y-auto p-5 space-y-5">
      <div className="grid grid-cols-4 gap-4">
        {[
          { svc: 'API Gateway', name: 'prahari-alb-gateway', lat: `${health.lat}ms`, st: health.status },
          { svc: 'WebSocket Hub', name: 'ws://prahari-events', lat: '27 Clients Connected', st: 'Active Stream' },
          { svc: 'PostgreSQL RDS', name: 'prahari-db-postgres', lat: 'v15.7 Engine', st: 'Healthy' },
          { svc: 'Redis Pub/Sub', name: 'prahari-redis-cache', lat: '40 events/sec', st: 'Healthy' },
          { svc: 'MQTT Broker', name: 'mqtt://prahari-broker', lat: '2,145 msg/min', st: 'Connected' },
          { svc: 'ECS Fargate Cluster', name: 'prahari-ecs-cluster', lat: '38% CPU • 56% RAM', st: '1 Task Running' },
          { svc: 'S3 Website Bucket', name: 'prahari-frontend-bucket', lat: 'HTTP 200 OK', st: 'Deployed' },
          { svc: 'CloudWatch Metrics', name: 'prahari-log-group', lat: '1,284 Signals Trace', st: 'Active Logging' },
        ].map(s => (
          <div key={s.svc} className="p-4 rounded-xl bg-white/[0.02] border border-white/[0.04]">
            <h3 className="text-xs font-bold text-white">{s.svc}</h3>
            <p className="text-[10px] text-zinc-500 mt-0.5">{s.name}</p>
            <div className="flex justify-between items-center mt-3 text-xs">
              <span className="text-zinc-400">{s.lat}</span>
              <span className="text-emerald-400 font-bold">{s.st}</span>
            </div>
          </div>
        ))}
      </div>
    </div>
  </div>
);

// ═══════════════════════════════════════════════════════════
// WORKSPACE 15: SETTINGS
// ═══════════════════════════════════════════════════════════
export const SettingsWorkspace: React.FC<{ session?: UserSession }> = ({ session }) => (
  <div className="h-full flex flex-col bg-[#09090b]">
    <Toolbar><span className="text-[11px] font-semibold text-zinc-300">SETTINGS & TENANT CONFIGURATION</span></Toolbar>
    <div className="flex-1 overflow-y-auto p-5 space-y-5">
      <div className="p-4 rounded-xl bg-white/[0.02] border border-white/[0.04] space-y-3">
        <h3 className="text-sm font-bold text-white">Tenant Organization Details</h3>
        <div className="grid grid-cols-2 gap-4 text-xs">
          <div><label className="text-[10px] text-zinc-500 uppercase">Organization Name</label><input disabled value={session?.orgName || 'Alpha Chemical Refinery Inc.'} className="w-full mt-1 px-3 py-2 rounded bg-white/[0.03] border border-white/[0.08] text-white" /></div>
          <div><label className="text-[10px] text-zinc-500 uppercase">User Role Scope</label><input disabled value={session?.role || 'Plant Manager'} className="w-full mt-1 px-3 py-2 rounded bg-white/[0.03] border border-white/[0.08] text-white" /></div>
          <div><label className="text-[10px] text-zinc-500 uppercase">AWS ALB API Endpoint</label><input disabled value="http://prahari-alb-hackathon-125438813.us-east-1.elb.amazonaws.com" className="w-full mt-1 px-3 py-2 rounded bg-white/[0.03] border border-white/[0.08] text-white font-mono" /></div>
          <div><label className="text-[10px] text-zinc-500 uppercase">Multi-LLM Fallback Model</label><input disabled value="OpenAI GPT-4o / Claude 3.5 / Bedrock" className="w-full mt-1 px-3 py-2 rounded bg-white/[0.03] border border-white/[0.08] text-white" /></div>
        </div>
      </div>
    </div>
  </div>
);
