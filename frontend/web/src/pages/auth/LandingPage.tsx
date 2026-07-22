import React from 'react';
import { Shield, Zap, Cpu, Eye, Activity, ChevronRight, CheckCircle2, ArrowRight, Server, Lock, Globe, Users } from 'lucide-react';

interface LandingPageProps {
  onGetStarted: () => void;
  onLogin: () => void;
  onDemoMode: (role: string) => void;
}

export const LandingPage: React.FC<LandingPageProps> = ({ onGetStarted, onLogin, onDemoMode }) => {
  return (
    <div className="min-h-screen bg-[#06080d] text-zinc-100 flex flex-col font-sans selection:bg-indigo-500 selection:text-white">
      {/* Top Header */}
      <header className="h-16 border-b border-white/[0.06] px-6 lg:px-12 flex items-center justify-between sticky top-0 bg-[#06080d]/80 backdrop-blur-md z-50">
        <div className="flex items-center gap-3 cursor-pointer" onClick={onGetStarted}>
          <div className="w-8 h-8 rounded-lg bg-indigo-600 flex items-center justify-center shadow-lg shadow-indigo-500/30">
            <Zap size={16} className="text-white" />
          </div>
          <span className="font-display text-lg font-bold text-white tracking-tight">PRAHARI</span>
          <span className="text-[10px] font-semibold px-2 py-0.5 rounded-full bg-indigo-500/10 text-indigo-400 border border-indigo-500/20">
            Enterprise OS v19.0
          </span>
        </div>

        <nav className="hidden md:flex items-center gap-8 text-sm font-medium text-zinc-400">
          <a href="#features" className="hover:text-white transition-colors">Platform</a>
          <a href="#architecture" className="hover:text-white transition-colors">Multi-Agent AI</a>
          <a href="#compliance" className="hover:text-white transition-colors">Compliance</a>
          <a href="#pricing" className="hover:text-white transition-colors">Enterprise</a>
        </nav>

        <div className="flex items-center gap-3">
          <button
            onClick={onLogin}
            className="px-4 py-2 rounded-lg text-sm font-medium text-zinc-300 hover:text-white hover:bg-white/[0.04] transition-colors"
          >
            Sign In
          </button>
          <button
            onClick={onGetStarted}
            className="px-4 py-2 rounded-lg text-sm font-semibold bg-indigo-600 hover:bg-indigo-500 text-white shadow-lg shadow-indigo-600/30 transition-all flex items-center gap-1.5"
          >
            Start Onboarding <ArrowRight size={14} />
          </button>
        </div>
      </header>

      {/* Hero Section */}
      <section className="py-20 px-6 lg:px-12 max-w-6xl mx-auto text-center space-y-8">
        <div className="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-white/[0.03] border border-white/[0.08] text-xs font-medium text-zinc-400">
          <span className="w-2 h-2 rounded-full bg-emerald-400 animate-pulse" />
          ISO 45001 & OSHA Compliant Industrial Safety Platform
        </div>

        <h1 className="font-display text-4xl sm:text-6xl font-extrabold text-white tracking-tight leading-[1.1]">
          The Industrial AI Safety <br />
          <span className="bg-gradient-to-r from-indigo-400 via-purple-300 to-emerald-400 bg-clip-text text-transparent">
            Operating System
          </span>
        </h1>

        <p className="text-lg text-zinc-400 max-w-2xl mx-auto leading-relaxed">
          Unify real-time plant telemetry, 10-agent AI supervision, computer vision PPE tracking, and predictive maintenance into a single enterprise control system.
        </p>

        <div className="flex flex-wrap items-center justify-center gap-4 pt-4">
          <button
            onClick={onGetStarted}
            className="px-6 py-3.5 rounded-xl bg-indigo-600 hover:bg-indigo-500 text-white font-semibold text-sm shadow-xl shadow-indigo-600/30 transition-all flex items-center gap-2"
          >
            Deploy Enterprise Tenant <ChevronRight size={16} />
          </button>
          <button
            onClick={() => onDemoMode('Plant Manager')}
            className="px-6 py-3.5 rounded-xl bg-white/[0.04] hover:bg-white/[0.08] border border-white/[0.08] text-zinc-200 font-semibold text-sm transition-all"
          >
            Launch Instant Demo Mode
          </button>
        </div>

        {/* Hero Preview Box */}
        <div className="pt-10 max-w-5xl mx-auto">
          <div className="rounded-2xl border border-white/[0.08] bg-zinc-900/60 p-2 shadow-2xl backdrop-blur-xl">
            <div className="rounded-xl bg-[#09090b] border border-white/[0.04] overflow-hidden p-6 text-left space-y-6">
              <div className="flex items-center justify-between border-b border-white/[0.06] pb-4">
                <div className="flex items-center gap-3">
                  <div className="w-3 h-3 rounded-full bg-red-500/80" />
                  <div className="w-3 h-3 rounded-full bg-yellow-500/80" />
                  <div className="w-3 h-3 rounded-full bg-green-500/80" />
                  <span className="text-xs text-zinc-500 ml-2 font-mono">prahari-command-center.enterprise.io</span>
                </div>
                <div className="flex items-center gap-2 text-xs text-emerald-400 font-medium">
                  <span className="w-1.5 h-1.5 rounded-full bg-emerald-400 animate-pulse" /> Live Telemetry Sync: 1,284 Signals
                </div>
              </div>

              <div className="grid grid-cols-4 gap-4">
                <div className="p-4 rounded-xl bg-white/[0.02] border border-white/[0.04]">
                  <p className="text-[11px] text-zinc-500 uppercase tracking-wider">Plant Health</p>
                  <p className="text-2xl font-bold text-emerald-400 mt-1">94.2%</p>
                  <p className="text-[10px] text-zinc-500 mt-1">Normal Operation</p>
                </div>
                <div className="p-4 rounded-xl bg-white/[0.02] border border-white/[0.04]">
                  <p className="text-[11px] text-zinc-500 uppercase tracking-wider">AI Agents Active</p>
                  <p className="text-2xl font-bold text-indigo-400 mt-1">10 / 10</p>
                  <p className="text-[10px] text-zinc-500 mt-1">Autonomous Reasoning</p>
                </div>
                <div className="p-4 rounded-xl bg-white/[0.02] border border-white/[0.04]">
                  <p className="text-[11px] text-zinc-500 uppercase tracking-wider">Edge Inference</p>
                  <p className="text-2xl font-bold text-white mt-1">30.0 FPS</p>
                  <p className="text-[10px] text-zinc-500 mt-1">11ms Latency</p>
                </div>
                <div className="p-4 rounded-xl bg-white/[0.02] border border-white/[0.04]">
                  <p className="text-[11px] text-zinc-500 uppercase tracking-wider">Active Permits</p>
                  <p className="text-2xl font-bold text-white mt-1">28 PTW</p>
                  <p className="text-[10px] text-zinc-500 mt-1">LOTO Verified</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Feature Grid */}
      <section id="features" className="py-20 px-6 lg:px-12 bg-white/[0.01] border-t border-white/[0.04]">
        <div className="max-w-6xl mx-auto space-y-12">
          <div className="text-center space-y-3">
            <h2 className="font-display text-3xl font-bold text-white">Engineered for Industrial Scale</h2>
            <p className="text-zinc-400 text-sm max-w-xl mx-auto">
              From continuous vibration monitoring to automated ISO compliance reports, PRAHARI delivers total plant safety visibility.
            </p>
          </div>

          <div className="grid md:grid-cols-3 gap-6">
            <div className="p-6 rounded-2xl bg-zinc-900/40 border border-white/[0.06] space-y-4">
              <div className="w-10 h-10 rounded-xl bg-indigo-500/10 flex items-center justify-center text-indigo-400">
                <Cpu size={20} />
              </div>
              <h3 className="text-lg font-semibold text-white">Multi-Agent AI Supervisor</h3>
              <p className="text-sm text-zinc-400 leading-relaxed">
                10 specialized domain agents collaborate synchronously to compute 5-Whys root causes, recalculate equipment RUL, and issue isolation locks.
              </p>
            </div>

            <div className="p-6 rounded-2xl bg-zinc-900/40 border border-white/[0.06] space-y-4">
              <div className="w-10 h-10 rounded-xl bg-emerald-500/10 flex items-center justify-center text-emerald-400">
                <Eye size={20} />
              </div>
              <h3 className="text-lg font-semibold text-white">YOLOv8 Edge Vision AI</h3>
              <p className="text-sm text-zinc-400 leading-relaxed">
                Real-time computer vision scans high-risk plant zones to enforce hardhat/PPE compliance and restrict unauthorized contractor access.
              </p>
            </div>

            <div className="p-6 rounded-2xl bg-zinc-900/40 border border-white/[0.06] space-y-4">
              <div className="w-10 h-10 rounded-xl bg-purple-500/10 flex items-center justify-center text-purple-400">
                <Activity size={20} />
              </div>
              <h3 className="text-lg font-semibold text-white">Physics-Informed Digital Twin</h3>
              <p className="text-sm text-zinc-400 leading-relaxed">
                Interactive CAD/SCADA twin schematic maps temperature, pressure, vibration, and failure modes across plant reactors, pumps, and boilers.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Role Quick Selector */}
      <section className="py-16 px-6 lg:px-12 border-t border-white/[0.04] bg-[#09090d]">
        <div className="max-w-4xl mx-auto text-center space-y-6">
          <h3 className="text-xl font-bold text-white">Experience PRAHARI by Persona</h3>
          <p className="text-sm text-zinc-400">Select a pre-configured role to immediately launch into the active command workspace:</p>
          <div className="flex flex-wrap items-center justify-center gap-3 pt-2">
            {['Plant Manager', 'Safety Officer', 'Reliability Engineer', 'Maintenance Engineer', 'Administrator'].map(role => (
              <button
                key={role}
                onClick={() => onDemoMode(role)}
                className="px-4 py-2 rounded-lg bg-zinc-800/80 hover:bg-indigo-600 hover:text-white border border-white/[0.08] text-xs font-medium text-zinc-300 transition-all flex items-center gap-2"
              >
                <Users size={14} /> {role} Demo
              </button>
            ))}
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="mt-auto border-t border-white/[0.06] py-8 px-6 lg:px-12 text-center text-xs text-zinc-500">
        <p>© 2026 PRAHARI Systems Inc. All rights reserved. Powered by AWS Cloud Infrastructure.</p>
      </footer>
    </div>
  );
};
