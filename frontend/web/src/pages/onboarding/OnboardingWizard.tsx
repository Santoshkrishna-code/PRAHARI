import React, { useState } from 'react';
import { Zap, CheckCircle2, ArrowRight, ArrowLeft, Building2, Factory, Database, Users, Brain, Sparkles } from 'lucide-react';
import { OrganizationSetup, UserSession } from '../../types/auth';

interface OnboardingWizardProps {
  onComplete: (setup: OrganizationSetup) => void;
  onSkip: () => void;
}

export const OnboardingWizard: React.FC<OnboardingWizardProps> = ({ onComplete, onSkip }) => {
  const [step, setStep] = useState(1);
  const [setup, setSetup] = useState<OrganizationSetup>({
    companyName: 'Alpha Chemical Refinery Inc.',
    industry: 'Oil & Gas / Chemical Processing',
    country: 'United States',
    timezone: 'UTC-05:00 (Eastern Time)',
    plantName: 'Plant Alpha (Gulf Coast)',
    location: 'Houston, TX',
    assetSource: 'IoT Edge Sensor CSV + MQTT Stream',
    teamEmail: 'safety-team@alpha-refinery.io',
    aiModel: 'OpenAI GPT-4o (Autonomous Safety Agent)',
  });

  const nextStep = () => {
    if (step < 5) setStep(s => s + 1);
    else onComplete(setup);
  };

  const prevStep = () => {
    if (step > 1) setStep(s => s - 1);
  };

  return (
    <div className="min-h-screen bg-[#06080d] text-zinc-100 flex flex-col justify-center items-center px-4 font-sans py-12 relative overflow-hidden">
      {/* Background Radial Glow */}
      <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[700px] h-[700px] bg-indigo-600/10 rounded-full blur-[160px] pointer-events-none" />

      {/* Onboarding Box */}
      <div className="w-full max-w-2xl bg-zinc-900/90 border border-white/[0.08] rounded-2xl p-8 shadow-2xl backdrop-blur-xl space-y-8 relative z-10">
        
        {/* Top Header */}
        <div className="flex items-center justify-between border-b border-white/[0.06] pb-6">
          <div className="flex items-center gap-3">
            <div className="w-9 h-9 rounded-xl bg-indigo-600 flex items-center justify-center shadow-lg shadow-indigo-500/30">
              <Zap size={18} className="text-white" />
            </div>
            <div>
              <h2 className="font-display text-lg font-bold text-white">PRAHARI Tenant Setup Wizard</h2>
              <p className="text-xs text-zinc-400">Step {step} of 5 — Configure your Enterprise AI Platform</p>
            </div>
          </div>
          <button
            onClick={onSkip}
            className="text-xs text-zinc-500 hover:text-zinc-300 transition-colors"
          >
            Skip to Command Center →
          </button>
        </div>

        {/* Progress Bar */}
        <div className="flex items-center justify-between gap-2">
          {[
            { id: 1, label: 'Organization', icon: <Building2 size={14} /> },
            { id: 2, label: 'Plant Site', icon: <Factory size={14} /> },
            { id: 3, label: 'Data Sources', icon: <Database size={14} /> },
            { id: 4, label: 'Team Roles', icon: <Users size={14} /> },
            { id: 5, label: 'AI Supervisor', icon: <Brain size={14} /> },
          ].map((s) => (
            <div
              key={s.id}
              className={`flex-1 flex flex-col items-center gap-1.5 cursor-pointer ${
                step >= s.id ? 'text-indigo-400 font-semibold' : 'text-zinc-600'
              }`}
              onClick={() => setStep(s.id)}
            >
              <div
                className={`w-8 h-8 rounded-xl flex items-center justify-center text-xs transition-all ${
                  step === s.id
                    ? 'bg-indigo-600 text-white shadow-lg shadow-indigo-500/40 ring-2 ring-indigo-400/40'
                    : step > s.id
                    ? 'bg-emerald-500/20 text-emerald-400 border border-emerald-500/30'
                    : 'bg-zinc-800 text-zinc-500'
                }`}
              >
                {step > s.id ? <CheckCircle2 size={16} /> : s.icon}
              </div>
              <span className="text-[10px] hidden sm:block">{s.label}</span>
            </div>
          ))}
        </div>

        {/* Step Content */}
        <div className="min-h-[240px] flex flex-col justify-center">
          {step === 1 && (
            <div className="space-y-4">
              <div className="space-y-1">
                <h3 className="text-base font-bold text-white">Create Your Organization</h3>
                <p className="text-xs text-zinc-400">Specify your company identity and regional operational timezone.</p>
              </div>
              <div className="space-y-3">
                <div>
                  <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1">
                    Company Name
                  </label>
                  <input
                    type="text"
                    value={setup.companyName}
                    onChange={e => setSetup({ ...setup, companyName: e.target.value })}
                    className="w-full px-3.5 py-2.5 rounded-xl bg-white/[0.03] border border-white/[0.08] text-sm text-white focus:outline-none focus:border-indigo-500"
                  />
                </div>
                <div className="grid grid-cols-2 gap-3">
                  <div>
                    <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1">
                      Industry Sector
                    </label>
                    <input
                      type="text"
                      value={setup.industry}
                      onChange={e => setSetup({ ...setup, industry: e.target.value })}
                      className="w-full px-3.5 py-2.5 rounded-xl bg-white/[0.03] border border-white/[0.08] text-sm text-white"
                    />
                  </div>
                  <div>
                    <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1">
                      Time Zone
                    </label>
                    <input
                      type="text"
                      value={setup.timezone}
                      onChange={e => setSetup({ ...setup, timezone: e.target.value })}
                      className="w-full px-3.5 py-2.5 rounded-xl bg-white/[0.03] border border-white/[0.08] text-sm text-white"
                    />
                  </div>
                </div>
              </div>
            </div>
          )}

          {step === 2 && (
            <div className="space-y-4">
              <div className="space-y-1">
                <h3 className="text-base font-bold text-white">Configure First Industrial Site</h3>
                <p className="text-xs text-zinc-400">Define the primary refinery or manufacturing facility for monitoring.</p>
              </div>
              <div className="space-y-3">
                <div>
                  <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1">
                    Plant Facility Name
                  </label>
                  <input
                    type="text"
                    value={setup.plantName}
                    onChange={e => setSetup({ ...setup, plantName: e.target.value })}
                    className="w-full px-3.5 py-2.5 rounded-xl bg-white/[0.03] border border-white/[0.08] text-sm text-white"
                  />
                </div>
                <div>
                  <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1">
                    Geographic Location
                  </label>
                  <input
                    type="text"
                    value={setup.location}
                    onChange={e => setSetup({ ...setup, location: e.target.value })}
                    className="w-full px-3.5 py-2.5 rounded-xl bg-white/[0.03] border border-white/[0.08] text-sm text-white"
                  />
                </div>
              </div>
            </div>
          )}

          {step === 3 && (
            <div className="space-y-4">
              <div className="space-y-1">
                <h3 className="text-base font-bold text-white">Connect Plant Data & Telemetry Sources</h3>
                <p className="text-xs text-zinc-400">Select how live sensor telemetry and assets will stream into PRAHARI.</p>
              </div>
              <div className="grid grid-cols-2 gap-3">
                {[
                  'IoT Edge Sensor CSV + MQTT Stream',
                  'REST API Gateway (/api/v1/telemetry)',
                  'SCADA / OPC-UA Pipeline',
                  'IBM Maximo Asset Database Direct Sync',
                ].map(opt => (
                  <div
                    key={opt}
                    onClick={() => setSetup({ ...setup, assetSource: opt })}
                    className={`p-3.5 rounded-xl border cursor-pointer transition-all text-xs font-medium ${
                      setup.assetSource === opt
                        ? 'bg-indigo-600/15 border-indigo-500 text-indigo-300'
                        : 'bg-white/[0.02] border-white/[0.06] text-zinc-400 hover:text-white'
                    }`}
                  >
                    {opt}
                  </div>
                ))}
              </div>
            </div>
          )}

          {step === 4 && (
            <div className="space-y-4">
              <div className="space-y-1">
                <h3 className="text-base font-bold text-white">Invite Safety & Reliability Team</h3>
                <p className="text-xs text-zinc-400">Assign role-based access for plant managers, safety officers, and engineers.</p>
              </div>
              <div>
                <label className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider block mb-1">
                  Team Invitation Email(s)
                </label>
                <input
                  type="text"
                  value={setup.teamEmail}
                  onChange={e => setSetup({ ...setup, teamEmail: e.target.value })}
                  className="w-full px-3.5 py-2.5 rounded-xl bg-white/[0.03] border border-white/[0.08] text-sm text-white"
                />
              </div>
            </div>
          )}

          {step === 5 && (
            <div className="space-y-4">
              <div className="space-y-1">
                <h3 className="text-base font-bold text-white">Configure Multi-Agent AI Model</h3>
                <p className="text-xs text-zinc-400">Choose the foundation LLM for real-time safety reasoning and RCA execution.</p>
              </div>
              <div className="space-y-2">
                {[
                  'OpenAI GPT-4o (Autonomous Safety Agent)',
                  'Claude 3.5 Sonnet (High-Precision Risk Reasoning)',
                  'Gemini 1.5 Pro (Multimodal Vision & Telemetry)',
                  'Ollama / Llama3 (On-Premises Air-Gapped Engine)',
                ].map(m => (
                  <div
                    key={m}
                    onClick={() => setSetup({ ...setup, aiModel: m })}
                    className={`p-3 rounded-xl border cursor-pointer transition-all text-xs font-medium flex items-center justify-between ${
                      setup.aiModel === m
                        ? 'bg-indigo-600/15 border-indigo-500 text-indigo-300'
                        : 'bg-white/[0.02] border-white/[0.06] text-zinc-400 hover:text-white'
                    }`}
                  >
                    <span>{m}</span>
                    {setup.aiModel === m && <Sparkles size={14} className="text-indigo-400" />}
                  </div>
                ))}
              </div>
            </div>
          )}
        </div>

        {/* Wizard Footer Buttons */}
        <div className="flex items-center justify-between border-t border-white/[0.06] pt-6">
          <button
            onClick={prevStep}
            disabled={step === 1}
            className={`px-4 py-2 rounded-xl text-xs font-medium flex items-center gap-1.5 ${
              step === 1 ? 'opacity-40 cursor-not-allowed text-zinc-600' : 'text-zinc-300 hover:bg-white/[0.04]'
            }`}
          >
            <ArrowLeft size={14} /> Back
          </button>

          <button
            onClick={nextStep}
            className="px-6 py-2.5 rounded-xl bg-indigo-600 hover:bg-indigo-500 text-white font-semibold text-xs shadow-lg shadow-indigo-600/30 transition-all flex items-center gap-1.5"
          >
            {step === 5 ? 'Launch PRAHARI Platform' : 'Next Step'} <ArrowRight size={14} />
          </button>
        </div>
      </div>
    </div>
  );
};
