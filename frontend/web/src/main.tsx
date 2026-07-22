import './index.css';
import React, { useState, useEffect } from 'react';
import { createRoot } from 'react-dom/client';
import {
  LayoutDashboard, Activity, BarChart3, FileText, Radio, Box, Package, Wrench,
  Shield, AlertTriangle, ClipboardCheck, Target, Brain, Cpu, Eye, Server,
  Settings, ChevronRight, ChevronDown, Search, Bell, User, Zap, ArrowUpRight,
  ArrowDownRight, Minus, Clock, CheckCircle2, XCircle, AlertCircle, Sparkles,
  MessageSquare, Play, Circle, TrendingUp, ChevronLeft, Command, Filter,
  Download, Plus, RefreshCw, Maximize2, Layers, Thermometer, Gauge, BarChart,
  PieChart, Calendar, MapPin, Hash, ArrowRight, ExternalLink, Wifi, Database,
  Lock, Unlock, FileSearch, GitBranch, MoreHorizontal, LogOut, UserCheck, X
} from 'lucide-react';

import { PageId } from './types';
import { UserSession, UserRole, OrganizationSetup } from './types/auth';
import { useHealth, useTelemetry } from './hooks/useTelemetry';
import { NAV_CONFIG } from './components/layout/NavConfig';
import { Toolbar, ToolBtn, ToolSep, Chart, Metric } from './components/common/CommonUI';

import { LandingPage } from './pages/auth/LandingPage';
import { LoginPage } from './pages/auth/LoginPage';
import { OnboardingWizard } from './pages/onboarding/OnboardingWizard';
import { CreateAssetModal, ReportIncidentModal } from './components/common/Modals';
import { GuidedTour } from './components/common/GuidedTour';
import { AIDrawer } from './components/common/AIDrawer';
import { EntityInspector } from './components/common/EntityInspector';
import { eventBus, EnterpriseEvent } from './services/eventBus';

// Import workspace views
import {
  CommandCenter,
  OperationsCenter,
  IndustrialTwin,
  AICommandCenter,
  IncidentsWorkspace,
  VisionIntelligence,
  AgentOrchestration,
  AssetsWorkspace,
  PlatformOps,
  OpsIntelligence,
  PermitsWorkspace,
  MaintenanceWorkspace,
  RiskWorkspace,
  InspectionsWorkspace,
  ExecutiveInsights,
  SettingsWorkspace
} from './main-workspaces';

const API = (window as any).__PRAHARI_API__ || '';

// ═══════════════════════════════════════════════════════════
// MAIN CONNECTED REAL-TIME ENTERPRISE AI SAFETY OPERATING SYSTEM
// ═══════════════════════════════════════════════════════════
const App: React.FC = () => {
  const [appView, setAppView] = useState<'landing' | 'login' | 'onboarding' | 'workspace'>('landing');
  const [page, setPage] = useState<PageId>('command-center');
  const [navOpen, setNavOpen] = useState(true);
  const [cmdOpen, setCmdOpen] = useState(false);
  const [tourActive, setTourActive] = useState(false);
  const [aiDrawerOpen, setAiDrawerOpen] = useState(false);
  const [inspectedEntityId, setInspectedEntityId] = useState<string | null>(null);

  // Proactive AI Alert Toast
  const [proactiveToast, setProactiveToast] = useState<{ title: string; desc: string } | null>(null);

  // Selected asset for Digital Twin context preservation
  const [selectedAssetId, setSelectedAssetId] = useState<string>('pump');

  // Modals state
  const [assetModalOpen, setAssetModalOpen] = useState(false);
  const [incidentModalOpen, setIncidentModalOpen] = useState(false);

  // Event stream state
  const [unreadCount, setUnreadCount] = useState(2);
  const [latestEvent, setLatestEvent] = useState<EnterpriseEvent | null>(null);

  // User session state
  const [userSession, setUserSession] = useState<UserSession>({
    id: 'usr-9901',
    name: 'Santosh Krishna',
    email: 'santosh@prahari-industrial.io',
    role: 'Plant Manager',
    orgName: 'Alpha Chemical Refinery Inc.',
    plantName: 'Plant Alpha (Gulf Coast)',
    avatar: 'SK',
    token: 'jwt-prahari-enterprise-token-9901',
  });

  const health = useHealth();
  const tele = useTelemetry();

  // Unified Event Bus subscription for real-time proactive push
  useEffect(() => {
    const unsubscribe = eventBus.subscribe((evt) => {
      setLatestEvent(evt);
      setUnreadCount(c => c + 1);

      if (evt.aiDecision && evt.severity === 'Critical') {
        setProactiveToast({
          title: `⚠️ Proactive AI Recommendation (${evt.asset || 'Plant Anomaly'})`,
          desc: evt.aiDecision,
        });
      }
    });
    return () => unsubscribe();
  }, []);

  useEffect(() => {
    const h = (e: KeyboardEvent) => {
      if ((e.metaKey || e.ctrlKey) && e.key === 'k') { e.preventDefault(); setCmdOpen(v => !v); }
      if ((e.metaKey || e.ctrlKey) && e.key === 'j') { e.preventDefault(); setAiDrawerOpen(v => !v); }
      if (e.key === 'Escape') { setCmdOpen(false); setAiDrawerOpen(false); setInspectedEntityId(null); }
    };
    window.addEventListener('keydown', h);
    return () => window.removeEventListener('keydown', h);
  }, []);

  const handleLoginSuccess = (user: { name?: string; email?: string; role?: UserRole; orgName?: string; plantName?: string }) => {
    const safeName = user?.name || 'Plant Manager';
    const initials = safeName.split(' ').filter(Boolean).map(n => n[0]).join('').slice(0, 2) || 'PM';
    setUserSession({
      ...userSession,
      name: safeName,
      email: user?.email || 'admin@prahari-industrial.io',
      role: user?.role || 'Plant Manager',
      orgName: user?.orgName || 'Alpha Chemical Refinery Inc.',
      plantName: user?.plantName || 'Plant Alpha (Gulf Coast)',
      avatar: initials,
    });
    setAppView('workspace');
    setTourActive(true);
  };

  const handleDemoLaunch = (role: string) => {
    setUserSession({
      ...userSession,
      role: role as UserRole,
      name: `${role} Lead`,
    });
    setAppView('workspace');
  };

  const sc = health.status === 'Operational' ? 'text-emerald-400' : 'text-amber-400';

  if (appView === 'landing') {
    return (
      <LandingPage
        onGetStarted={() => setAppView('onboarding')}
        onLogin={() => setAppView('login')}
        onDemoMode={(role: string) => {
          // Direct demo clicks to authentication login view
          setAppView('login');
        }}
      />
    );
  }

  if (appView === 'login') {
    return (
      <LoginPage
        onLoginSuccess={handleLoginSuccess}
        onNavigateRegister={() => setAppView('onboarding')}
        onNavigateLanding={() => setAppView('landing')}
      />
    );
  }

  if (appView === 'onboarding') {
    return (
      <OnboardingWizard
        onComplete={(setup: OrganizationSetup) => {
          setUserSession({
            ...userSession,
            orgName: setup.companyName,
            plantName: setup.plantName,
          });
          setAppView('workspace');
          setTourActive(true);
        }}
        onSkip={() => setAppView('workspace')}
      />
    );
  }

  const handleNavigateToAsset = (assetId: string) => {
    setSelectedAssetId(assetId);
    setPage('industrial-twin');
  };

  const handleNavigateToPage = (targetPage: PageId) => {
    setPage(targetPage);
  };

  const handleGenerateWorkOrder = (asset: string, desc: string) => {
    eventBus.emit({
      eventId: `evt-wo-${Date.now()}`,
      timestamp: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }),
      category: 'Maintenance',
      source: 'AI Supervisor',
      asset: asset,
      plant: 'Plant Alpha (Gulf Coast)',
      severity: 'Info',
      correlationId: `corr-wo-${asset}`,
      message: `Generated Work Order for ${asset}: ${desc}`,
      aiDecision: 'Auto-dispatched Work Order to CMMS Maintenance Queue',
      priority: 1,
    });
    setProactiveToast({
      title: `✅ Work Order Dispatched (${asset})`,
      desc: `Work Order created and dispatched to CMMS Maintenance Queue. RUL recalculated.`,
    });
  };

  const renderWorkspace = () => {
    switch (page) {
      case 'command-center': return <CommandCenter tele={tele} onReportIncident={() => setIncidentModalOpen(true)} onNavigateToAsset={handleNavigateToAsset} onNavigateToPage={handleNavigateToPage} onGenerateWorkOrder={handleGenerateWorkOrder} />;
      case 'operations': return <OperationsCenter tele={tele} onNavigateToAsset={handleNavigateToAsset} />;
      case 'industrial-twin': return <IndustrialTwin tele={tele} initialSelectedAssetId={selectedAssetId} />;
      case 'ai-command': return <AICommandCenter onGenerateWorkOrder={handleGenerateWorkOrder} />;
      case 'incidents': return <IncidentsWorkspace onReportIncident={() => setIncidentModalOpen(true)} />;
      case 'vision-intel': return <VisionIntelligence />;
      case 'agent-orch': return <AgentOrchestration />;
      case 'assets': return <AssetsWorkspace tele={tele} onAddAsset={() => setAssetModalOpen(true)} onNavigateToAsset={handleNavigateToAsset} />;
      case 'platform-ops': return <PlatformOps health={health} />;
      case 'ops-intelligence': return <OpsIntelligence tele={tele} onNavigateToPage={handleNavigateToPage} onGenerateWorkOrder={handleGenerateWorkOrder} />;
      case 'permits': return <PermitsWorkspace />;
      case 'maintenance': return <MaintenanceWorkspace />;
      case 'risk': return <RiskWorkspace onNavigateToAsset={handleNavigateToAsset} />;
      case 'inspections': return <InspectionsWorkspace />;
      case 'executive': return <ExecutiveInsights />;
      case 'settings': return <SettingsWorkspace session={userSession} />;
      default: return <CommandCenter tele={tele} onReportIncident={() => setIncidentModalOpen(true)} onNavigateToAsset={handleNavigateToAsset} onNavigateToPage={handleNavigateToPage} onGenerateWorkOrder={handleGenerateWorkOrder} />;
    }
  };

  return (
    <div className="h-screen flex flex-col bg-[#09090b] text-zinc-200 overflow-hidden font-sans relative">
      {/* Proactive AI Alert Toast Banner */}
      {proactiveToast && (
        <div className="absolute top-14 right-6 z-50 max-w-md bg-amber-950/95 border border-amber-500/40 rounded-2xl p-4 shadow-2xl backdrop-blur-md">
          <div className="flex items-start justify-between">
            <div className="flex items-center gap-2">
              <Sparkles size={16} className="text-amber-400" />
              <h4 className="font-bold text-amber-300 text-xs">{proactiveToast.title}</h4>
            </div>
            <button onClick={() => setProactiveToast(null)} className="text-zinc-500 hover:text-white">
              <X size={14} />
            </button>
          </div>
          <p className="text-[12px] text-zinc-200 mt-1.5 leading-relaxed">{proactiveToast.desc}</p>
          <div className="mt-3 flex flex-wrap gap-2">
            <button
              onClick={() => { handleGenerateWorkOrder('PUMP-P102', 'Bearing race replacement'); }}
              className="px-2.5 py-1 rounded bg-amber-500 text-black font-bold text-[10px] hover:bg-amber-400"
            >
              Generate Work Order
            </button>
            <button
              onClick={() => { handleNavigateToAsset('pump'); setProactiveToast(null); }}
              className="px-2.5 py-1 rounded bg-white/[0.08] hover:bg-white/[0.15] text-white font-semibold text-[10px]"
            >
              View Digital Twin
            </button>
            <button
              onClick={() => { setPage('ai-command'); setProactiveToast(null); }}
              className="px-2.5 py-1 rounded bg-indigo-600 hover:bg-indigo-500 text-white font-semibold text-[10px]"
            >
              Open AI Command Center
            </button>
          </div>
        </div>
      )}

      {/* Guided Tour Banner */}
      {tourActive && (
        <GuidedTour
          onNavigate={(p) => setPage(p)}
          onDismiss={() => setTourActive(false)}
        />
      )}

      {/* Universal Ask AI Slide-out Drawer */}
      <AIDrawer
        isOpen={aiDrawerOpen}
        onClose={() => setAiDrawerOpen(false)}
        currentPage={page}
        selectedEntityId={inspectedEntityId || undefined}
      />

      {/* Cross-Module Entity Inspector Drawer */}
      <EntityInspector
        entityId={inspectedEntityId}
        onClose={() => setInspectedEntityId(null)}
        onSelectEntity={(id) => setInspectedEntityId(id)}
      />

      {/* Top Header */}
      <header className="h-11 border-b border-white/[0.06] px-3 flex items-center justify-between shrink-0 z-30 bg-[#09090b]">
        <div className="flex items-center gap-2.5">
          <div
            className="w-7 h-7 rounded-lg bg-indigo-600 flex items-center justify-center cursor-pointer shadow-lg shadow-indigo-500/20"
            onClick={() => setAppView('landing')}
          >
            <Zap size={13} className="text-white" />
          </div>
          <span className="font-display text-[14px] font-bold text-white tracking-tight cursor-pointer" onClick={() => setAppView('landing')}>
            PRAHARI
          </span>
          <span className="text-[10px] px-1.5 py-0.5 rounded bg-zinc-800/80 text-zinc-400 border border-white/[0.06]">
            {userSession.orgName}
          </span>
          <span className="text-[10px] px-1.5 py-0.5 rounded bg-indigo-500/10 text-indigo-300 border border-indigo-500/20 font-medium">
            Role: {userSession.role}
          </span>
        </div>

        <div className="flex items-center gap-2">
          {/* Universal Ask AI Button */}
          <button
            onClick={() => setAiDrawerOpen(v => !v)}
            className="flex items-center gap-1.5 px-3 py-1 rounded-lg bg-indigo-600 hover:bg-indigo-500 text-white font-semibold text-[11px] shadow-lg shadow-indigo-600/30 transition-all"
          >
            <Brain size={13} />
            <span>Ask AI</span>
            <kbd className="text-[9px] px-1 rounded bg-indigo-700/60 text-white ml-1">⌘J</kbd>
          </button>

          <button
            onClick={() => setCmdOpen(true)}
            className="flex items-center gap-2 px-2.5 py-1 rounded-lg bg-zinc-900/80 border border-white/[0.06] text-zinc-500 hover:text-zinc-300 transition-colors text-[12px]"
          >
            <Search size={13} />
            <span>Search</span>
            <kbd className="text-[9px] px-1 py-0.5 rounded bg-zinc-800 text-zinc-600 ml-2">⌘K</kbd>
          </button>

          <button
            onClick={() => setInspectedEntityId('PUMP-P102')}
            className="px-2 py-1 rounded-lg bg-zinc-900 border border-white/[0.06] text-zinc-400 hover:text-white text-[10px] font-medium flex items-center gap-1"
          >
            <GitBranch size={12} className="text-indigo-400" /> Graph
          </button>

          <button
            onClick={() => setUnreadCount(0)}
            className="w-7 h-7 rounded-lg hover:bg-zinc-800 flex items-center justify-center text-zinc-500 hover:text-zinc-300 transition-colors relative"
          >
            <Bell size={15} />
            {unreadCount > 0 && (
              <span className="absolute top-1 right-1 w-2 h-2 rounded-full bg-indigo-500 animate-pulse" />
            )}
          </button>

          <button
            onClick={() => setAppView('login')}
            title="Sign Out"
            className="w-7 h-7 rounded-lg hover:bg-zinc-800 flex items-center justify-center text-zinc-500 hover:text-zinc-300 transition-colors"
          >
            <LogOut size={14} />
          </button>

          <div className="w-7 h-7 rounded-lg bg-indigo-600 flex items-center justify-center text-[10px] font-bold text-white shadow">
            {userSession.avatar}
          </div>
        </div>
      </header>

      <div className="flex-1 flex overflow-hidden">
        {/* Nav Rail */}
        <nav className={`${navOpen ? 'w-52' : 'w-12'} border-r border-white/[0.06] flex flex-col shrink-0 transition-all duration-200 overflow-hidden bg-[#09090b]`}>
          <div className="flex-1 overflow-y-auto py-2 px-1.5">
            {NAV_CONFIG.map((s, si) => (
              <div key={si} className={si > 0 ? 'mt-3.5' : ''}>
                {navOpen && <p className="text-[9px] font-semibold text-zinc-600 uppercase tracking-[0.15em] px-2 mb-1">{s.label}</p>}
                {s.items.map(item => {
                  const active = page === item.id;
                  return (
                    <button
                      key={item.id}
                      onClick={() => setPage(item.id)}
                      className={`w-full flex items-center gap-2 px-2 py-1.5 rounded-md text-left transition-all duration-100 group relative mb-px
                        ${active ? 'bg-white/[0.06] text-white' : 'text-zinc-500 hover:text-zinc-300 hover:bg-white/[0.03]'}`}
                    >
                      {active && <div className="absolute left-0 top-1/2 -translate-y-1/2 w-0.5 h-3.5 rounded-full bg-indigo-500" />}
                      <span className={`shrink-0 ${active ? 'text-indigo-400' : item.accent ? 'text-indigo-500/60' : 'text-zinc-600 group-hover:text-zinc-400'} transition-colors`}>{item.icon}</span>
                      {navOpen && (
                        <>
                          <span className="text-[12px] font-medium flex-1 truncate">{item.label}</span>
                          {item.badge && <span className="text-[9px] font-medium px-1.5 py-0.5 rounded-full bg-zinc-800 text-zinc-400">{item.badge}</span>}
                        </>
                      )}
                    </button>
                  );
                })}
              </div>
            ))}
          </div>
          <button
            onClick={() => setNavOpen(v => !v)}
            className="h-8 border-t border-white/[0.06] flex items-center justify-center text-zinc-600 hover:text-zinc-400 transition-colors shrink-0"
          >
            {navOpen ? <ChevronLeft size={14} /> : <ChevronRight size={14} />}
          </button>
        </nav>

        {/* Active Workspace */}
        <main className="flex-1 overflow-hidden" key={page}>
          <div className="h-full page-enter">{renderWorkspace()}</div>
        </main>
      </div>

      {/* Status Bar */}
      <footer className="h-6 border-t border-white/[0.06] px-3 flex items-center justify-between text-[10px] text-zinc-600 shrink-0 bg-[#09090b]">
        <div className="flex items-center gap-3">
          <span className={`flex items-center gap-1 ${sc}`}>
            <span className={`w-1.5 h-1.5 rounded-full ${health.status === 'Operational' ? 'bg-emerald-400 animate-pulse' : 'bg-amber-400'}`} />
            {health.status}
          </span>
          <span className="text-zinc-400">API Latency: <strong className="text-emerald-400">148ms</strong></span>
          <span className="text-zinc-400">Sync Interval: <strong className="text-zinc-300">1.5s</strong></span>
          <span className="text-zinc-400">WebSocket: <strong className="text-emerald-400">Connected</strong></span>
          {latestEvent && (
            <span className="text-zinc-400 truncate max-w-sm">
              Stream Event: <strong>{latestEvent.source}</strong> — {latestEvent.message}
            </span>
          )}
        </div>
        <div className="flex items-center gap-3">
          <span>Tenant: {userSession.orgName}</span>
          <span>AWS 727533783159</span>
          <span>US-EAST-1</span>
          <span>ISO 45001</span>
        </div>
      </footer>

      {/* Command Palette */}
      {cmdOpen && (
        <div className="fixed inset-0 z-50 bg-black/60 backdrop-blur-sm flex items-start justify-center pt-[12vh]" onClick={() => setCmdOpen(false)}>
          <div className="w-[480px] rounded-2xl bg-zinc-900 border border-white/[0.08] shadow-2xl overflow-hidden" onClick={e => e.stopPropagation()}>
            <div className="px-3.5 py-2.5 border-b border-white/[0.06] flex items-center gap-2">
              <Search size={15} className="text-zinc-500" />
              <input autoFocus className="w-full bg-transparent text-sm text-white placeholder-zinc-500 focus:outline-none" placeholder="Search workspaces, assets, agents..." />
            </div>
            <div className="p-1.5 max-h-72 overflow-y-auto">
              {NAV_CONFIG.flatMap(s => s.items).map(item => (
                <button
                  key={item.id}
                  onClick={() => { setPage(item.id); setCmdOpen(false); }}
                  className="w-full flex items-center gap-2.5 px-3 py-2 rounded-lg hover:bg-zinc-800 transition-colors group text-left"
                >
                  <span className="text-zinc-500 group-hover:text-zinc-300">{item.icon}</span>
                  <span className="text-[13px] text-zinc-300 group-hover:text-white">{item.label}</span>
                </button>
              ))}
            </div>
          </div>
        </div>
      )}

      {/* CRUD Modals */}
      <CreateAssetModal
        isOpen={assetModalOpen}
        onClose={() => setAssetModalOpen(false)}
        onCreate={(asset) => {
          eventBus.emit({
            category: 'Asset',
            source: 'Assets Workspace',
            asset: asset.name,
            severity: 'Info',
            correlationId: `corr-${asset.name}`,
            message: `Registered asset ${asset.name} in location ${asset.loc}`,
            priority: 1,
          });
        }}
      />

      <ReportIncidentModal
        isOpen={incidentModalOpen}
        onClose={() => setIncidentModalOpen(false)}
        onReport={(incident) => {
          eventBus.emit({
            category: 'Incident',
            source: 'Incidents Workspace',
            asset: incident.asset,
            severity: 'Warning',
            correlationId: incident.id,
            message: `Dispatched incident ${incident.id}: ${incident.title}`,
            priority: 3,
          });
        }}
      />
    </div>
  );
};

interface ErrorBoundaryState {
  hasError: boolean;
  error: Error | null;
}

class ErrorBoundary extends React.Component<{ children: React.ReactNode }, ErrorBoundaryState> {
  constructor(props: { children: React.ReactNode }) {
    super(props);
    this.state = { hasError: false, error: null };
  }

  static getDerivedStateFromError(error: Error): ErrorBoundaryState {
    return { hasError: true, error };
  }

  componentDidCatch(error: Error, errorInfo: React.ErrorInfo) {
    console.error('PRAHARI Unhandled Workspace Error:', error, errorInfo);
  }

  render() {
    if (this.state.hasError) {
      return (
        <div className="min-h-screen bg-[#06080d] text-zinc-100 flex flex-col justify-center items-center p-6 font-sans">
          <div className="max-w-md w-full bg-zinc-900 border border-white/[0.08] rounded-2xl p-6 shadow-2xl text-center space-y-4">
            <div className="w-12 h-12 rounded-xl bg-red-500/10 text-red-400 flex items-center justify-center mx-auto">
              <AlertTriangle size={24} />
            </div>
            <h2 className="text-lg font-bold text-white">Workspace Exception Intercepted</h2>
            <p className="text-xs text-zinc-400 leading-relaxed">
              PRAHARI Error Boundary caught an exception during workspace render.
            </p>
            <div className="p-3 rounded-lg bg-black/50 border border-white/[0.06] text-left text-[11px] font-mono text-red-300 max-h-32 overflow-y-auto">
              {this.state.error?.message || 'Unknown render error'}
            </div>
            <button
              onClick={() => { this.setState({ hasError: false, error: null }); window.location.reload(); }}
              className="px-4 py-2 rounded-xl bg-indigo-600 hover:bg-indigo-500 text-white font-bold text-xs shadow-lg shadow-indigo-600/30 transition-all w-full"
            >
              Reload Operational Control Center
            </button>
          </div>
        </div>
      );
    }
    return this.props.children;
  }
}

createRoot(document.getElementById('root')!).render(
  <ErrorBoundary>
    <App />
  </ErrorBoundary>
);
