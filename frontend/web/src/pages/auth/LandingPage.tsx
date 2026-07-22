import React, { useState } from 'react';
import {
  Shield, Zap, Cpu, Eye, Activity, ChevronRight, CheckCircle2, ArrowRight, Server, Lock, Globe, Users, Sparkles, BarChart3, Radio, RefreshCw, FileText, Check, AlertTriangle, Layers
} from 'lucide-react';

interface LandingPageProps {
  onGetStarted: () => void;
  onLogin: () => void;
  onDemoMode: (role: string) => void;
}

export const LandingPage: React.FC<LandingPageProps> = ({ onGetStarted, onLogin, onDemoMode }) => {
  const [activeTab, setActiveTab] = useState<'command' | 'twin' | 'vision' | 'compliance'>('command');

  return (
    <div className="min-h-screen bg-[#06080d] text-zinc-100 flex flex-col font-sans selection:bg-indigo-500 selection:text-white">
      {/* Top Navigation Bar */}
      <header className="h-16 border-b border-white/[0.06] px-6 lg:px-12 flex items-center justify-between sticky top-0 bg-[#06080d]/80 backdrop-blur-md z-50">
        <div className="flex items-center gap-3 cursor-pointer" onClick={onGetStarted}>
          <div className="w-8 h-8 rounded-lg bg-indigo-600 flex items-center justify-center shadow-lg shadow-indigo-500/30">
            <Zap size={16} className="text-white" />
          </div>
          <span className="font-display text-lg font-bold text-white tracking-tight">PRAHARI</span>
          <span className="text-[10px] font-semibold px-2 py-0.5 rounded-full bg-indigo-500/10 text-indigo-400 border border-indigo-500/20">
            Enterprise OS v20.0
          </span>
        </div>

        <nav className="hidden md:flex items-center gap-8 text-xs font-medium text-zinc-400">
          <a href="#features" className="hover:text-white transition-colors">Capabilities</a>
          <a href="#architecture" className="hover:text-white transition-colors">Architecture</a>
          <a href="#compliance" className="hover:text-white transition-colors">ISO 45001</a>
          <a href="#personas" className="hover:text-white transition-colors">Role Demos</a>
        </nav>

        <div className="flex items-center gap-3">
          <button
            onClick={onLogin}
            className="px-4 py-2 rounded-lg text-xs font-medium text-zinc-300 hover:text-white hover:bg-white/[0.04] transition-colors"
          >
            Sign In
          </button>
          <button
            onClick={onGetStarted}
            className="px-4 py-2 rounded-lg text-xs font-semibold bg-indigo-600 hover:bg-indigo-500 text-white shadow-lg shadow-indigo-600/30 transition-all flex items-center gap-1.5"
          >
            Deploy Enterprise Tenant <ArrowRight size={14} />
          </button>
        </div>
      </header>

      {/* Hero Section */}
      <section className="py-16 px-6 lg:px-12 max-w-6xl mx-auto text-center space-y-8">
        <div className="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-white/[0.03] border border-white/[0.08] text-xs font-medium text-zinc-400">
          <span className="w-2 h-2 rounded-full bg-emerald-400 animate-pulse" />
          Real-Time Industrial AI Safety & Asset Operating System
        </div>

        <h1 className="font-display text-4xl sm:text-6xl font-extrabold text-white tracking-tight leading-[1.1]">
          The Enterprise AI Platform for <br />
          <span className="bg-gradient-to-r from-indigo-400 via-purple-300 to-emerald-400 bg-clip-text text-transparent">
            Zero-Incident Industrial Operations
          </span>
        </h1>

        <p className="text-base sm:text-lg text-zinc-400 max-w-2xl mx-auto leading-relaxed">
          Unify OPC-UA telemetry, PINN digital twin physics, YOLOv8 edge vision, and 10-agent autonomous supervisors into a truthful single source of operational intelligence.
        </p>

        <div className="flex flex-wrap items-center justify-center gap-4 pt-2">
          <button
            onClick={() => onDemoMode('Plant Manager')}
            className="px-6 py-3.5 rounded-xl bg-indigo-600 hover:bg-indigo-500 text-white font-semibold text-sm shadow-xl shadow-indigo-600/30 transition-all flex items-center gap-2"
          >
            Launch Command Center (Plant Manager Shift B) <ChevronRight size={16} />
          </button>
          <button
            onClick={onGetStarted}
            className="px-6 py-3.5 rounded-xl bg-white/[0.04] hover:bg-white/[0.08] border border-white/[0.08] text-zinc-200 font-semibold text-sm transition-all"
          >
            Start 5-Step Enterprise Setup
          </button>
        </div>

        {/* Live Metrics Proof Bar */}
        <div className="grid grid-cols-4 gap-4 max-w-4xl mx-auto pt-6 text-center">
          <div className="p-3 rounded-xl bg-white/[0.02] border border-white/[0.04]">
            <p className="text-[10px] text-zinc-500 uppercase tracking-wider">AWS ALB Latency</p>
            <p className="text-lg font-bold text-emerald-400">148 ms</p>
          </div>
          <div className="p-3 rounded-xl bg-white/[0.02] border border-white/[0.04]">
            <p className="text-[10px] text-zinc-500 uppercase tracking-wider">Signals Monitored</p>
            <p className="text-lg font-bold text-white">1,284 / sec</p>
          </div>
          <div className="p-3 rounded-xl bg-white/[0.02] border border-white/[0.04]">
            <p className="text-[10px] text-zinc-500 uppercase tracking-wider">AI Confidence</p>
            <p className="text-lg font-bold text-indigo-400">97.8%</p>
          </div>
          <div className="p-3 rounded-xl bg-white/[0.02] border border-white/[0.04]">
            <p className="text-[10px] text-zinc-500 uppercase tracking-wider">OSHA Recordables</p>
            <p className="text-lg font-bold text-emerald-400">0.00 TRIR</p>
          </div>
        </div>

        {/* Hero Product Command Center Preview */}
        <div className="pt-6 max-w-5xl mx-auto">
          <div className="rounded-2xl border border-white/[0.08] bg-zinc-900/60 p-2 shadow-2xl backdrop-blur-xl">
            <div className="rounded-xl bg-[#09090b] border border-white/[0.04] overflow-hidden p-5 text-left space-y-4">
              <div className="flex items-center justify-between border-b border-white/[0.06] pb-3">
                <div className="flex items-center gap-3">
                  <div className="w-3 h-3 rounded-full bg-red-500/80" />
                  <div className="w-3 h-3 rounded-full bg-yellow-500/80" />
                  <div className="w-3 h-3 rounded-full bg-emerald-500/80" />
                  <span className="text-xs text-zinc-400 font-mono">prahari-os.plant-alpha.enterprise.io</span>
                </div>
                <div className="flex items-center gap-2 text-xs text-emerald-400 font-medium">
                  <span className="w-1.5 h-1.5 rounded-full bg-emerald-400 animate-pulse" /> Live Telemetry Engine Connected
                </div>
              </div>

              <div className="grid grid-cols-4 gap-3">
                <div className="p-3.5 rounded-xl bg-white/[0.02] border border-white/[0.04]">
                  <p className="text-[10px] text-zinc-500 uppercase tracking-wider">Safety Index</p>
                  <p className="text-xl font-bold text-emerald-400 mt-0.5">94.2/100</p>
                  <p className="text-[10px] text-zinc-600 mt-1">↑ +2.1% from yesterday</p>
                </div>
                <div className="p-3.5 rounded-xl bg-white/[0.02] border border-white/[0.04]">
                  <p className="text-[10px] text-zinc-500 uppercase tracking-wider">Vibration (P-102)</p>
                  <p className="text-xl font-bold text-amber-400 mt-0.5">11.8 mm/s</p>
                  <p className="text-[10px] text-zinc-600 mt-1">RUL: 18 days remaining</p>
                </div>
                <div className="p-3.5 rounded-xl bg-white/[0.02] border border-white/[0.04]">
                  <p className="text-[10px] text-zinc-500 uppercase tracking-wider">YOLOv8 Vision</p>
                  <p className="text-xl font-bold text-white mt-0.5">29.8 FPS</p>
                  <p className="text-[10px] text-zinc-600 mt-1">CAM-002 Zone B scan</p>
                </div>
                <div className="p-3.5 rounded-xl bg-white/[0.02] border border-white/[0.04]">
                  <p className="text-[10px] text-zinc-500 uppercase tracking-wider">LOTO Isolation</p>
                  <p className="text-xl font-bold text-emerald-400 mt-0.5">100% Sync</p>
                  <p className="text-[10px] text-zinc-600 mt-1">28 active permits</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Interactive Platform Feature Tabs */}
      <section id="features" className="py-16 px-6 lg:px-12 bg-white/[0.01] border-t border-white/[0.04]">
        <div className="max-w-6xl mx-auto space-y-8">
          <div className="text-center space-y-3">
            <h2 className="font-display text-2xl sm:text-3xl font-bold text-white">Complete Industrial Operational Platform</h2>
            <p className="text-zinc-400 text-xs sm:text-sm max-w-xl mx-auto">
              Built for refinery managers, EHS compliance leads, and reliability engineers to eliminate downtime and safety hazards.
            </p>
          </div>

          {/* Tab Selector */}
          <div className="flex justify-center gap-2 border-b border-white/[0.06] pb-3">
            {[
              { id: 'command', label: 'Command Center' },
              { id: 'twin', label: 'Digital Twin Physics' },
              { id: 'vision', label: 'YOLOv8 Edge Vision' },
              { id: 'compliance', label: 'ISO 45001 Audit' },
            ].map(t => (
              <button
                key={t.id}
                onClick={() => setActiveTab(t.id as any)}
                className={`px-4 py-2 rounded-lg text-xs font-semibold transition-all ${activeTab === t.id ? 'bg-indigo-600 text-white shadow-lg shadow-indigo-600/30' : 'text-zinc-400 hover:text-white hover:bg-white/[0.03]'}`}
              >
                {t.label}
              </button>
            ))}
          </div>

          {/* Tab Content Box */}
          <div className="p-6 rounded-2xl bg-zinc-900/60 border border-white/[0.06] text-xs text-zinc-300">
            {activeTab === 'command' && (
              <div className="space-y-3">
                <h3 className="text-base font-bold text-white">Central Operations & Telemetry Command Center</h3>
                <p className="text-zinc-400 leading-relaxed">
                  Real-time SCADA signal monitoring with executive BI metrics, proactive AI recommendations, and continuous 30-day historical trend graphs for industrial pumps, heat exchangers, and reactors.
                </p>
              </div>
            )}
            {activeTab === 'twin' && (
              <div className="space-y-3">
                <h3 className="text-base font-bold text-white">PINN Physics-Informed Digital Twin</h3>
                <p className="text-zinc-400 leading-relaxed">
                  Interactive CAD/SCADA twin schematic evaluating bearing outer race wear probabilities, RUL forecasts, and thermal breakdown limits under actual line load.
                </p>
              </div>
            )}
            {activeTab === 'vision' && (
              <div className="space-y-3">
                <h3 className="text-base font-bold text-white">YOLOv8 Edge Vision AI Intelligence</h3>
                <p className="text-zinc-400 leading-relaxed">
                  Jetson AGX edge inferencing scanning high-risk plant zones to detect hardhat violations, restricted area intrusions, and contractor badge expiration.
                </p>
              </div>
            )}
            {activeTab === 'compliance' && (
              <div className="space-y-3">
                <h3 className="text-base font-bold text-white">Automated ISO 45001 & OSHA Audit Center</h3>
                <p className="text-zinc-400 leading-relaxed">
                  Instant evidence document verification, zone safety score heatmaps, 5-Whys root cause trace generation, and automated LOTO isolation verification.
                </p>
              </div>
            )}
          </div>
        </div>
      </section>

      {/* Role Quick Selector */}
      <section id="personas" className="py-16 px-6 lg:px-12 border-t border-white/[0.04] bg-[#09090d]">
        <div className="max-w-4xl mx-auto text-center space-y-6">
          <h3 className="text-xl font-bold text-white">Launch PRAHARI Persona Demos</h3>
          <p className="text-xs text-zinc-400">Select any role below to launch directly into that specialized enterprise workspace:</p>
          <div className="flex flex-wrap items-center justify-center gap-3 pt-2">
            {[
              { role: 'Plant Manager', desc: 'Executive BI & Fleet Health' },
              { role: 'Safety Officer', desc: 'ISO 45001 & Incidents' },
              { role: 'Reliability Engineer', desc: 'Digital Twin & RUL' },
              { role: 'Maintenance Engineer', desc: 'CMMS Work Orders' },
              { role: 'Administrator', desc: 'AWS Ops & Access' },
            ].map(item => (
              <button
                key={item.role}
                onClick={() => onDemoMode(item.role)}
                className="px-4 py-2.5 rounded-xl bg-zinc-800/80 hover:bg-indigo-600 hover:text-white border border-white/[0.08] text-xs font-medium text-zinc-300 transition-all flex items-center gap-2 group"
              >
                <Users size={14} className="text-indigo-400 group-hover:text-white" />
                <div className="text-left">
                  <span className="block font-bold">{item.role}</span>
                  <span className="text-[9px] text-zinc-500 group-hover:text-indigo-200">{item.desc}</span>
                </div>
              </button>
            ))}
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="mt-auto border-t border-white/[0.06] py-6 px-6 lg:px-12 text-center text-xs text-zinc-500">
        <p>© 2026 PRAHARI Systems Inc. All rights reserved. Powered by AWS Cloud & Go Telemetry Engine.</p>
      </footer>
    </div>
  );
};
