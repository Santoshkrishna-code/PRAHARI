import React, { useState } from 'react';
import { createRoot } from 'react-dom/client';
import { ThemeProvider } from '@prahari/theme/ThemeProvider.tsx';
import { PermissionsProvider } from '@prahari/utils/permissions.tsx';
import { AuthProvider } from '@prahari/utils/auth.tsx';
import { Dashboard } from './screens/dashboard/Dashboard.tsx';
import { Permits } from './screens/permits/Permits.tsx';
import { Inspections } from './screens/inspections/Inspections.tsx';
import { Incidents } from './screens/incidents/Incidents.tsx';
import { AICopilot } from './screens/ai/AICopilot.tsx';

import '../../frontend/platform/apps/playground/index.css';

const AppContent: React.FC = () => {
  const [activeTab, setActiveTab] = useState<string>('dashboard');

  const tabs = [
    { id: 'dashboard', label: 'Home' },
    { id: 'permits', label: 'Permits' },
    { id: 'inspections', label: 'Inspections' },
    { id: 'incidents', label: 'Report' },
    { id: 'ai', label: 'Copilot' }
  ];

  const renderScreen = () => {
    switch (activeTab) {
      case 'dashboard':
        return <Dashboard onNavigate={setActiveTab} />;
      case 'permits':
        return <Permits />;
      case 'inspections':
        return <Inspections />;
      case 'incidents':
        return <Incidents />;
      case 'ai':
        return <AICopilot />;
      default:
        return <Dashboard onNavigate={setActiveTab} />;
    }
  };

  return (
    <div className="flex flex-col min-h-screen bg-background text-text">
      {/* Header bar */}
      <header className="h-14 border-b border-border bg-surface flex items-center justify-between px-4 flex-shrink-0">
        <span className="text-sm font-bold tracking-tight uppercase">PRAHARI Mobile iOS</span>
      </header>

      {/* Main container */}
      <main className="flex-1 overflow-y-auto bg-background pb-16">
        {renderScreen()}
      </main>

      {/* Bottom Navigation Tabs */}
      <nav className="fixed bottom-0 left-0 right-0 h-16 border-t border-border bg-surface flex items-center justify-around z-40">
        {tabs.map((tab) => {
          const isActive = activeTab === tab.id;
          return (
            <button
              key={tab.id}
              onClick={() => setActiveTab(tab.id)}
              className={`flex flex-col items-center justify-center gap-0.5 text-xs font-semibold py-1.5 px-2.5 rounded transition-colors ${
                isActive ? 'text-primary bg-border/20' : 'text-text/75 hover:text-text'
              }`}
            >
              <span>{tab.label}</span>
            </button>
          );
        })}
      </nav>
    </div>
  );
};

const container = document.getElementById('root');
if (container) {
  const root = createRoot(container);
  root.render(
    <React.StrictMode>
      <ThemeProvider>
        <AuthProvider>
          <PermissionsProvider>
            <AppContent />
          </PermissionsProvider>
        </AuthProvider>
      </ThemeProvider>
    </React.StrictMode>
  );
}
export default AppContent;
