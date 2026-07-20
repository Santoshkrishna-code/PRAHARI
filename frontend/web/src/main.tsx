import React, { useState } from 'react';
import { createRoot } from 'react-dom/client';
import { ThemeProvider } from '@prahari/theme/ThemeProvider.tsx';
import { PermissionsProvider } from '@prahari/utils/permissions.tsx';
import { DashboardLayout } from '@prahari/layouts/DashboardLayout.tsx';
import { Sidebar } from '@prahari/navigation/Sidebar.tsx';

// Modules
import { Dashboard } from './modules/dashboard/Dashboard.tsx';
import { Permits } from './modules/permits/Permits.tsx';
import { Incidents } from './modules/incidents/Incidents.tsx';
import { AICopilot } from './modules/ai/AICopilot.tsx';
import { VisionConsole } from './modules/vision/VisionConsole.tsx';
import { TwinCanvas } from './modules/digitaltwin/TwinCanvas.tsx';

import '../../platform/apps/playground/index.css';

const AppContent: React.FC = () => {
  const [activeTab, setActiveTab] = useState<string>('dashboard');

  const sidebarItems = [
    { label: 'Control Center Dashboard', href: 'dashboard' },
    { label: 'Safe Work Permits', href: 'permits' },
    { label: 'Incidents Registry', href: 'incidents' },
    { label: 'AI Copilot Help', href: 'ai' },
    { label: 'Vision Stream Grid', href: 'vision' },
    { label: 'Digital Twin layout', href: 'digitaltwin' }
  ];

  const handleNavigate = (href: string) => {
    setActiveTab(href);
  };

  const sidebar = (
    <Sidebar items={sidebarItems} currentPath={activeTab} onNavigate={handleNavigate} />
  );

  const header = (
    <div className="flex justify-between items-center w-full">
      <span className="text-base font-bold text-text">PRAHARI Control Panel Portal</span>
      <div className="flex gap-2.5 items-center">
        <span className="text-xs font-semibold text-text/80 bg-surface border border-border rounded px-2.5 py-1">
          Role: System Administrator
        </span>
      </div>
    </div>
  );

  const renderModule = () => {
    switch (activeTab) {
      case 'dashboard':
        return <Dashboard />;
      case 'permits':
        return <Permits />;
      case 'incidents':
        return <Incidents />;
      case 'ai':
        return <AICopilot />;
      case 'vision':
        return <VisionConsole />;
      case 'digitaltwin':
        return <TwinCanvas />;
      default:
        return <Dashboard />;
    }
  };

  return (
    <DashboardLayout sidebar={sidebar} header={header}>
      {renderModule()}
    </DashboardLayout>
  );
};

const container = document.getElementById('root');
if (container) {
  const root = createRoot(container);
  root.render(
    <React.StrictMode>
      <ThemeProvider>
        <PermissionsProvider>
          <AppContent />
        </PermissionsProvider>
      </ThemeProvider>
    </React.StrictMode>
  );
}
