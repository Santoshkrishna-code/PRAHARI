import React, { useState, useEffect } from 'react';
import {
  LayoutDashboard, Activity, BarChart3, FileText, Radio, Box, Package, Wrench,
  Shield, AlertTriangle, ClipboardCheck, Target, Brain, Cpu, Eye, Server,
  Settings, ChevronRight, ChevronDown, Search, Bell, User, Zap, ArrowUpRight,
  ArrowDownRight, Minus, Clock, CheckCircle2, XCircle, AlertCircle, Sparkles,
  MessageSquare, Play, Circle, TrendingUp, ChevronLeft, Command, Filter,
  Download, Plus, RefreshCw, Maximize2, Layers, Thermometer, Gauge, BarChart,
  PieChart, Calendar, MapPin, Hash, ArrowRight, ExternalLink, Wifi, Database,
  Lock, Unlock, FileSearch, GitBranch, MoreHorizontal, LogOut, UserCheck,
  SlidersHorizontal, Rewind, SkipForward, Users, HelpCircle
} from 'lucide-react';

import { TelemetryPoint, HealthStatus, Asset, Incident } from './types';
import { UserSession } from './types/auth';
import { Toolbar, ToolBtn, ToolSep, Chart, Metric } from './components/common/CommonUI';

// ═══════════════════════════════════════════════════════════
// WORKSPACE 1: COMMAND CENTER
// ═══════════════════════════════════════════════════════════
export const CommandCenter: React.FC<{ tele: TelemetryPoint[]; onReportIncident?: () => void }> = ({ tele, onReportIncident }) => {
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
        <ToolBtn><RefreshCw size={12} /> Refresh</ToolBtn>
        <ToolBtn><Maximize2 size={12} /></ToolBtn>
      </Toolbar>

      <div className="flex-1 overflow-y-auto p-5 space-y-5">
        {/* Executive Greeting Header */}
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

        {/* 6 Key Stat Cards */}
        <div className="grid grid-cols-6 gap-3">
          {[
            { l: 'Safety Index', v: '94.2', u: '/100', c: 'text-emerald-400', sub: '+2.1% vs yesterday' },
            { l: 'Incident Rate (TRIR)', v: '0.18', u: '', c: 'text-emerald-400', sub: 'Zero OSHA recordables' },
            { l: 'Active Risks', v: '2', u: 'High', c: 'text-amber-400', sub: 'DC-101 Loop' },
            { l: 'Inspection Pass', v: '98.4', u: '%', c: 'text-emerald-400', sub: '122 audited' },
            { l: 'Asset Health', v: '91.2', u: '%', c: 'text-white', sub: '47 monitored' },
            { l: 'Permit Compliance', v: '100', u: '%', c: 'text-emerald-400', sub: '28 active LOTO' },
          ].map(k => (
            <div key={k.l} className="p-3.5 rounded-xl bg-white/[0.02] border border-white/[0.04]">
              <p className="text-[9px] text-zinc-500 uppercase tracking-wider mb-1 font-semibold">{k.l}</p>
              <p className={`text-xl font-bold ${k.c}`}>{k.v}<span className="text-xs text-zinc-500 font-normal ml-0.5">{k.u}</span></p>
              <p className="text-[10px] text-zinc-600 mt-1">{k.sub}</p>
            </div>
          ))}
        </div>

        {/* 3-Column Layout */}
        <div className="grid grid-cols-3 gap-5">
          {/* Column 1: AI Recommendations */}
          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <div className="flex items-center gap-2">
                <Sparkles size={14} className="text-indigo-400" />
                <span className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider">AI Supervisor Actions</span>
              </div>
              <span className="text-[10px] text-zinc-600">3 Pending</span>
            </div>
            {[
              { pri: 'high', title: 'Replace bearing assembly — Pump P-102', summary: 'Vibration velocity (11.8 mm/s) trending toward 14.0 mm/s limit. RUL: 18 days. Bearing wear: 72%.', agent: 'Maintenance Agent', confidence: 94, time: '2m' },
              { pri: 'medium', title: 'Audit Zone B contractor badges', summary: 'Contractor C-4412 detected with expired badge near reactor. Gate B access auto-revoked.', agent: 'Contractor Agent', confidence: 92, time: '8m' },
              { pri: 'low', title: 'Schedule lubrication — Compressor C-03', summary: 'PM interval threshold reached in 6 days. Auto-escalation configured in CMMS.', agent: 'Maintenance Agent', confidence: 96, time: '14m' },
            ].map((r, i) => (
              <div key={i} className="p-3.5 rounded-xl bg-white/[0.02] hover:bg-white/[0.04] border border-white/[0.04] transition-colors cursor-pointer group">
                <div className="flex items-start gap-2.5">
                  <span className={`w-1.5 h-1.5 rounded-full mt-2 shrink-0 ${r.pri === 'high' ? 'bg-amber-400' : r.pri === 'medium' ? 'bg-indigo-400' : 'bg-zinc-600'}`} />
                  <div className="flex-1 min-w-0">
                    <p className="text-[13px] font-medium text-zinc-200 group-hover:text-white transition-colors">{r.title}</p>
                    <p className="text-[12px] text-zinc-500 mt-1 leading-relaxed">{r.summary}</p>
                    <div className="flex items-center gap-3 mt-2 text-[10px] text-zinc-600">
                      <span className="flex items-center gap-1"><Brain size={10} />{r.agent}</span>
                      <span>{r.confidence}% conf.</span>
                      <span>{r.time} ago</span>
                    </div>
                  </div>
                  <ChevronRight size={14} className="text-zinc-700 group-hover:text-zinc-400 mt-1 transition-colors" />
                </div>
              </div>
            ))}
          </div>

          {/* Column 2: Key Indicators & Gauges */}
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

          {/* Column 3: Operational Activity Stream */}
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
// WORKSPACE 2: OPERATIONS INTELLIGENCE (Rich Executive BI View)
// ═══════════════════════════════════════════════════════════
export const OpsIntelligence: React.FC<{ tele: TelemetryPoint[] }> = ({ tele }) => {
  const l = tele[tele.length - 1] || { vib: 11.8, temp: 94.1 };
  const vibTrend = tele.map(p => p.vib);
  const tempTrend = tele.map(p => p.temp);

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
        {/* Executive Greeting Header */}
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

        {/* 6 Key Executive KPIs */}
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

        {/* AI Executive Summary Briefing */}
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

        {/* 30-Day Historical Trend Charts */}
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

        {/* Risk Contribution Breakdown & Operational Action Center */}
        <div className="grid grid-cols-2 gap-5">
          {/* Risk Weight Breakdown */}
          <div className="p-4 rounded-xl bg-white/[0.015] border border-white/[0.04] space-y-3">
            <span className="text-sm font-bold text-white block">Plant Risk Contribution Breakdown</span>
            {[
              { category: 'Asset Health & RUL Degradation', weight: 34, color: 'bg-amber-500' },
              { category: 'Permit & LOTO Isolation Compliance', weight: 26, color: 'bg-indigo-500' },
              { category: 'Inspection & Checklist Pass Rate', weight: 18, color: 'bg-emerald-500' },
              { category: 'Incidents & OSHA Recordables', weight: 12, color: 'bg-purple-500' },
              { category: 'Environmental & Gas Thresholds', weight: 10, color: 'bg-cyan-500' },
            ].map(r => (
              <div key={r.category} className="space-y-1">
                <div className="flex justify-between text-xs">
                  <span className="text-zinc-300">{r.category}</span>
                  <span className="text-zinc-400 font-semibold">{r.weight}%</span>
                </div>
                <div className="w-full h-1.5 rounded-full bg-zinc-800">
                  <div className={`h-1.5 rounded-full ${r.color}`} style={{ width: `${r.weight}%` }} />
                </div>
              </div>
            ))}
          </div>

          {/* Executive Action Center */}
          <div className="p-4 rounded-xl bg-white/[0.015] border border-white/[0.04] space-y-3">
            <span className="text-sm font-bold text-white block">Action Center — Immediate Items</span>
            {[
              { item: 'Pump P-102 Bearing Work Order WO-7821', pri: 'Medium Risk', status: 'Overdue by 14d', action: 'Approve WO' },
              { item: 'Zone B Contractor Badge Audit', pri: 'Security', status: 'Completed', action: 'View Audit' },
              { item: 'Hot Work Permit PTW-8902 Gas Verification', pri: 'Safety', status: '28 Active LOTO', action: 'Review PTW' },
              { item: 'ISO 45001 Compliance Audit Export', pri: 'Compliance', status: 'Ready for PDF', action: 'Generate' },
            ].map((a, i) => (
              <div key={i} className="flex items-center justify-between p-2.5 rounded-lg bg-white/[0.02] border border-white/[0.04] text-xs">
                <div>
                  <span className="font-medium text-white block">{a.item}</span>
                  <span className="text-[10px] text-zinc-500">{a.pri} • {a.status}</span>
                </div>
                <button className="px-2.5 py-1 rounded bg-indigo-600 hover:bg-indigo-500 text-white font-semibold text-[10px]">
                  {a.action}
                </button>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 3: OPERATIONS CENTER
// ═══════════════════════════════════════════════════════════
export const OperationsCenter: React.FC<{ tele: TelemetryPoint[] }> = ({ tele }) => {
  const l = tele[tele.length - 1] || { vib: 11.8, temp: 94.1, psi: 242, kw: 330, flow: 84 };
  const [chartMode, setChartMode] = useState<'vib' | 'temp' | 'psi' | 'kw'>('vib');
  const chartConfig = {
    vib: { data: tele.map(p => p.vib), color: '#818cf8', threshold: 14, label: 'Vibration Velocity', unit: 'mm/s', alarm: 'ISO 10816 Limit @ 14.0' },
    temp: { data: tele.map(p => p.temp), color: '#f59e0b', threshold: 100, label: 'Bearing Temperature', unit: '°C', alarm: 'Alarm @ 100°C' },
    psi: { data: tele.map(p => p.psi), color: '#22d3ee', threshold: 260, label: 'Line Pressure', unit: 'PSI', alarm: 'High @ 260 PSI' },
    kw: { data: tele.map(p => +p.kw), color: '#a78bfa', threshold: 370, label: 'Power Draw', unit: 'kW', alarm: 'Peak @ 370 kW' },
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
        <span className="text-[10px] text-emerald-400 font-medium flex items-center gap-1"><span className="w-1.5 h-1.5 rounded-full bg-emerald-400 animate-pulse" /> Live Telemetry</span>
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
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 4: INDUSTRIAL TWIN
// ═══════════════════════════════════════════════════════════
export const IndustrialTwin: React.FC<{ tele: TelemetryPoint[] }> = ({ tele }) => {
  const l = tele[tele.length - 1] || { temp: 94.1, vib: 11.8 };
  const [selected, setSelected] = useState('pump');
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
        <span className="text-[10px] text-zinc-500">Reactor Complex B Schematic</span>
      </Toolbar>

      <div className="flex-1 flex overflow-hidden">
        <div className="flex-1 relative bg-[#0b0b0f]">
          <svg viewBox="0 0 800 460" className="w-full h-full" preserveAspectRatio="xMidYMid meet">
            {[[240,150,380,200],[460,200,580,140],[400,240,380,320],[180,190,160,290],[430,340,580,320]].map(([x1,y1,x2,y2], i) => (
              <line key={i} x1={x1} y1={y1} x2={x2} y2={y2} stroke="rgba(99,102,241,0.1)" strokeWidth="2" strokeDasharray="8,6" />
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
          <div className="w-72 border-l border-white/[0.04] p-4 bg-white/[0.01] space-y-4">
            <div>
              <p className="text-[10px] text-zinc-500 uppercase">Selected Equipment</p>
              <h2 className="text-base font-bold text-white mt-0.5">{sel.label}</h2>
              <span className={`inline-block mt-1 text-[10px] px-2 py-0.5 rounded-full ${sel.st === 'Warning' ? 'bg-amber-500/15 text-amber-400' : 'bg-emerald-500/10 text-emerald-400'}`}>{sel.st}</span>
            </div>
            <div className="grid grid-cols-2 gap-3">
              <Metric label="Health" value={`${sel.health}%`} small />
              <Metric label="RUL" value={sel.rul} small />
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 5: AI COMMAND CENTER
// ═══════════════════════════════════════════════════════════
export const AICommandCenter: React.FC = () => {
  const steps = [
    { text: 'Detected vibration anomaly on Pump P-102', status: 'done', detail: 'Reading 11.8 mm/s, trending toward ISO limit.' },
    { text: 'Queried CMMS maintenance history', status: 'done', detail: 'PM overdue by 14 days. Work order WO-7821 un-escalated.' },
    { text: 'Calculated bearing wear probability', status: 'done', detail: 'PINN digital twin predicts 72% wear probability. RUL: 18 days.' },
    { text: 'Verified PTW-8902 isolation lock', status: 'done', detail: 'Valve V-88 isolation locked. No conflicts.' },
    { text: 'Synthesizing corrective action plan (CAPA)', status: 'active', detail: '4 actions generated for maintenance crew.' },
  ];

  return (
    <div className="h-full flex flex-col bg-[#09090b]">
      <Toolbar>
        <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">AI COMMAND CENTER</span>
        <ToolSep />
        <span className="text-[10px] text-indigo-400 font-medium">10 Autonomous Agents Active</span>
      </Toolbar>

      <div className="flex-1 flex overflow-hidden">
        <div className="flex-1 p-5 space-y-4 overflow-y-auto">
          <div className="p-4 rounded-xl bg-indigo-600/10 border border-indigo-500/20">
            <h2 className="text-sm font-bold text-white">Active Reasoning: Pump P-102 Anomaly</h2>
            <p className="text-xs text-zinc-400 mt-1">Multi-agent supervisor is conducting 5-Whys RCA and evidence correlation.</p>
          </div>
          <div className="space-y-2">
            {steps.map((s, i) => (
              <div key={i} className="flex gap-3 py-2 px-3 rounded-lg bg-white/[0.02]">
                <CheckCircle2 size={16} className={s.status === 'done' ? 'text-emerald-400' : 'text-indigo-400'} />
                <div>
                  <p className="text-[12px] font-medium text-zinc-200">{s.text}</p>
                  <p className="text-[11px] text-zinc-500 mt-0.5">{s.detail}</p>
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
// WORKSPACE 6: INCIDENTS WORKSPACE
// ═══════════════════════════════════════════════════════════
export const IncidentsWorkspace: React.FC<{ onReportIncident?: () => void }> = ({ onReportIncident }) => (
  <div className="h-full flex flex-col bg-[#09090b]">
    <Toolbar>
      <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">INCIDENTS & HAZARDS</span>
      <div className="flex-1" />
      <ToolBtn onClick={onReportIncident} className="!bg-red-600/80 !text-white">
        <Plus size={12} /> Report Incident
      </ToolBtn>
    </Toolbar>
    <div className="flex-1 p-5 space-y-4 overflow-y-auto">
      <div className="p-4 rounded-xl bg-white/[0.02] border border-white/[0.04]">
        <div className="flex justify-between">
          <span className="text-xs font-mono text-zinc-500">INC-2026-0447</span>
          <span className="text-xs text-amber-400 font-medium">Under Investigation</span>
        </div>
        <h2 className="text-sm font-bold text-white mt-1">Pump P-102 Vibration Excursion</h2>
        <p className="text-xs text-zinc-400 mt-1">5-Whys Root Cause: PM overdue by 14 days led to bearing lubrication failure.</p>
      </div>
    </div>
  </div>
);

// ═══════════════════════════════════════════════════════════
// WORKSPACE 7: SAFE WORK PERMITS
// ═══════════════════════════════════════════════════════════
export const PermitsWorkspace: React.FC = () => (
  <div className="h-full flex flex-col bg-[#09090b]">
    <Toolbar>
      <span className="text-[11px] font-semibold text-zinc-300">SAFE WORK PERMITS</span>
      <ToolSep />
      <ToolBtn active>28 Active LOTO</ToolBtn>
      <div className="flex-1" />
      <ToolBtn className="!bg-indigo-600 !text-white"><Plus size={12} /> Issue PTW</ToolBtn>
    </Toolbar>
    <div className="flex-1 p-5 space-y-4 overflow-y-auto">
      <div className="h-8 flex items-center px-4 gap-2 bg-white/[0.01] border-b border-white/[0.04] text-[10px] text-zinc-500 uppercase tracking-wider">
        <span className="w-24">Permit ID</span><span className="flex-1">Description</span><span className="w-24">Type</span><span className="w-28">Isolation</span><span className="w-24">Gas Test</span><span className="w-20">Status</span>
      </div>
      {[
        { id: 'PTW-8902', desc: 'Hot Work — Tank T-204', type: 'Hot Work', iso: 'V-88 LOCKED', gas: 'O₂ 20.9%', st: 'Approved' },
        { id: 'PTW-8903', desc: 'Confined Entry — Vessel V-109', type: 'Confined', iso: 'LINE BLIND', gas: 'Pending', st: 'Hold' },
        { id: 'PTW-8904', desc: 'Electrical — Panel MCC-7B', type: 'Electrical', iso: 'RACKED OUT', gas: 'N/A', st: 'Approved' },
        { id: 'PTW-8905', desc: 'Excavation — Pipe Trench B', type: 'Excavation', iso: 'N/A', gas: 'N/A', st: 'Pending' },
      ].map(p => (
        <div key={p.id} className="h-10 flex items-center px-4 gap-2 text-[12px] hover:bg-white/[0.02] cursor-pointer transition-colors border-b border-white/[0.02]">
          <span className="w-24 text-indigo-400 font-medium font-mono text-[11px]">{p.id}</span>
          <span className="flex-1 text-zinc-300">{p.desc}</span>
          <span className="w-24 text-zinc-500">{p.type}</span>
          <span className="w-28 text-zinc-500">{p.iso}</span>
          <span className="w-24 text-zinc-500">{p.gas}</span>
          <span className="w-20"><span className={`text-[10px] px-1.5 py-0.5 rounded-full font-medium ${p.st === 'Approved' ? 'bg-emerald-500/10 text-emerald-400' : p.st === 'Hold' ? 'bg-amber-500/15 text-amber-400' : 'bg-zinc-800 text-zinc-400'}`}>{p.st}</span></span>
        </div>
      ))}
    </div>
  </div>
);

// ═══════════════════════════════════════════════════════════
// WORKSPACE 8: MAINTENANCE
// ═══════════════════════════════════════════════════════════
export const MaintenanceWorkspace: React.FC = () => (
  <div className="h-full flex flex-col bg-[#09090b]">
    <Toolbar><span className="text-[11px] font-semibold text-zinc-300">MAINTENANCE WORKFLOW</span></Toolbar>
    <div className="flex-1 p-5 space-y-4 overflow-y-auto">
      <div className="h-8 flex items-center px-4 gap-2 bg-white/[0.01] border-b border-white/[0.04] text-[10px] text-zinc-500 uppercase tracking-wider">
        <span className="w-24">WO ID</span><span className="flex-1">Description</span><span className="w-24">Asset</span><span className="w-16">Priority</span><span className="w-16">RUL</span><span className="w-20">Status</span>
      </div>
      {[
        { id: 'WO-7821', desc: 'Bearing replacement and lubrication service', asset: 'P-102', pri: 'Critical', rul: '18d', st: 'Overdue' },
        { id: 'WO-7822', desc: 'Vibration probe recalibration', asset: 'P-102', pri: 'High', rul: '18d', st: 'Assigned' },
        { id: 'WO-7823', desc: 'Quarterly compressor inspection', asset: 'C-03', pri: 'Medium', rul: '84d', st: 'Scheduled' },
        { id: 'WO-7824', desc: 'Boiler tube thickness measurement', asset: 'Boiler A', pri: 'Medium', rul: '71d', st: 'Scheduled' },
      ].map(w => (
        <div key={w.id} className="h-10 flex items-center px-4 gap-2 text-[12px] hover:bg-white/[0.02] cursor-pointer transition-colors border-b border-white/[0.02]">
          <span className="w-24 text-indigo-400 font-medium font-mono text-[11px]">{w.id}</span>
          <span className="flex-1 text-zinc-300">{w.desc}</span>
          <span className="w-24 text-zinc-500">{w.asset}</span>
          <span className="w-16"><span className={`text-[10px] font-medium ${w.pri === 'Critical' ? 'text-red-400' : 'text-amber-400'}`}>{w.pri}</span></span>
          <span className="w-16 text-amber-400">{w.rul}</span>
          <span className="w-20"><span className={`text-[10px] px-1.5 py-0.5 rounded-full font-medium ${w.st === 'Overdue' ? 'bg-red-500/15 text-red-400' : 'bg-indigo-500/15 text-indigo-400'}`}>{w.st}</span></span>
        </div>
      ))}
    </div>
  </div>
);

// ═══════════════════════════════════════════════════════════
// WORKSPACE 9: RISK ASSESSMENT
// ═══════════════════════════════════════════════════════════
export const RiskWorkspace: React.FC = () => (
  <div className="h-full flex flex-col bg-[#09090b]">
    <Toolbar><span className="text-[11px] font-semibold text-zinc-300">RISK ASSESSMENT</span></Toolbar>
    <div className="p-5 flex gap-6">
      <div className="flex-1">
        <p className="text-[11px] text-zinc-400 font-medium mb-3">5×5 Industrial Risk Assessment Matrix</p>
        <div className="grid grid-cols-6 gap-px bg-white/[0.04] rounded-xl overflow-hidden">
          <div className="bg-[#09090b] p-2" />{['Rare', 'Unlikely', 'Possible', 'Likely', 'Almost Certain'].map(h => <div key={h} className="bg-[#0d0d12] p-2 text-[9px] text-zinc-500 text-center">{h}</div>)}
          {['Catastrophic', 'Major', 'Moderate', 'Minor', 'Insignificant'].map((sev, si) => (
            <React.Fragment key={sev}>
              <div className="bg-[#0d0d12] p-2 text-[9px] text-zinc-500 text-right">{sev}</div>
              {[1,2,3,4,5].map(li => {
                const risk = (5 - si) * li;
                const isHighlighted = si === 1 && li === 3;
                return (
                  <div key={li} className={`p-2 text-center text-[11px] font-medium ${risk > 15 ? 'bg-red-500/10 text-red-400' : risk > 10 ? 'bg-amber-500/10 text-amber-400' : 'bg-emerald-500/5 text-emerald-400/60'} ${isHighlighted ? 'ring-1 ring-indigo-500' : ''}`}>
                    {risk}{isHighlighted && <span className="block text-[8px] text-indigo-400">P-102</span>}
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
// OTHER WORKSPACES
// ═══════════════════════════════════════════════════════════
export const VisionIntelligence: React.FC = () => (
  <div className="h-full flex flex-col bg-[#09090b]"><Toolbar><span className="text-[11px] font-semibold text-zinc-300">VISION INTELLIGENCE</span></Toolbar><div className="p-5 text-xs text-zinc-400">4-Camera YOLOv8 Inference Grid Operational</div></div>
);

export const AgentOrchestration: React.FC = () => (
  <div className="h-full flex flex-col bg-[#09090b]"><Toolbar><span className="text-[11px] font-semibold text-zinc-300">AGENT ORCHESTRATION</span></Toolbar><div className="p-5 text-xs text-zinc-400">10-Agent Execution Trace & Memory DAG</div></div>
);

export const AssetsWorkspace: React.FC<{ tele: TelemetryPoint[]; onAddAsset?: () => void }> = ({ onAddAsset }) => (
  <div className="h-full flex flex-col bg-[#09090b]">
    <Toolbar>
      <span className="text-[11px] font-semibold text-zinc-300">ASSETS WORKSPACE</span>
      <div className="flex-1" />
      <ToolBtn onClick={onAddAsset} className="!bg-indigo-600 !text-white"><Plus size={12} /> Add Asset</ToolBtn>
    </Toolbar>
    <div className="p-5 text-xs text-zinc-400">IBM Maximo Asset Management Grid (47 Monitored Assets)</div>
  </div>
);

export const PlatformOps: React.FC<{ health: HealthStatus }> = ({ health }) => (
  <div className="h-full flex flex-col bg-[#09090b]"><Toolbar><span className="text-[11px] font-semibold text-zinc-300">PLATFORM OPERATIONS</span></Toolbar><div className="p-5 text-xs text-zinc-400">AWS Infrastructure Observability: {health.status} ({health.lat}ms)</div></div>
);

export const InspectionsWorkspace: React.FC = () => (
  <div className="h-full flex flex-col bg-[#09090b]"><Toolbar><span className="text-[11px] font-semibold text-zinc-300">INSPECTIONS</span></Toolbar><div className="p-5 text-xs text-zinc-400">Compliance Audit Checklists</div></div>
);

export const ExecutiveInsights: React.FC = () => (
  <div className="h-full flex flex-col bg-[#09090b]"><Toolbar><span className="text-[11px] font-semibold text-zinc-300">EXECUTIVE INSIGHTS</span></Toolbar><div className="p-5 text-xs text-zinc-400">ISO 45001 & OSHA Audit Reports</div></div>
);

export const SettingsWorkspace: React.FC<{ session?: UserSession }> = ({ session }) => (
  <div className="h-full flex flex-col bg-[#09090b]">
    <Toolbar><span className="text-[11px] font-semibold text-zinc-300">SETTINGS</span></Toolbar>
    <div className="p-5 text-xs text-zinc-400">Organization: {session?.orgName || 'Alpha Chemical Refinery Inc.'} ({session?.role})</div>
  </div>
);
