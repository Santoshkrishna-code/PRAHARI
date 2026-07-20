import React, { useState } from 'react';
import { createRoot } from 'react-dom/client';
import { z } from 'zod';
import { ThemeProvider, useTheme } from '../../packages/theme/ThemeProvider.tsx';
import { Button } from '../../packages/components/Button.tsx';
import { Badge } from '../../packages/components/Badge.tsx';
import { Switch } from '../../packages/components/Switch.tsx';
import { Spinner } from '../../packages/components/Spinner.tsx';
import { Card } from '../../packages/components/Card.tsx';
import { KPICard } from '../../packages/components/KPICard.tsx';
import { DataTable } from '../../packages/components/DataTable.tsx';
import { Modal } from '../../packages/components/Modal.tsx';
import { Toast } from '../../packages/components/Toast.tsx';
import { DynamicForm } from '../../packages/forms/DynamicForm.tsx';
import { DashboardLayout } from '../../packages/layouts/DashboardLayout.tsx';
import { Sidebar } from '../../packages/navigation/Sidebar.tsx';
import { CommandPalette } from '../../packages/navigation/CommandPalette.tsx';
import { EChartWrapper } from '../../packages/charts/EChartWrapper.tsx';

import './index.css';

const formSchema = z.object({
  username: z.string().min(3, 'Username must be at least 3 characters'),
  email: z.string().email('Please enter a valid email address')
});

const formFields = [
  { name: 'username', label: 'Operator Username', placeholder: 'Enter username' },
  { name: 'email', label: 'Operator Email', placeholder: 'Enter email address' }
];

const tableColumns = [
  { header: 'ID', accessor: (item: any) => item.id },
  { header: 'Device / Area', accessor: (item: any) => item.name },
  { header: 'Status', accessor: (item: any) => <Badge variant={item.status === 'ONLINE' ? 'success' : 'danger'}>{item.status}</Badge> }
];

const tableData = [
  { id: 'cam-001', name: 'Refinery Distillation Column Cam', status: 'ONLINE' },
  { id: 'cam-002', name: 'Hazard Gas Area Sensor Feed', status: 'OFFLINE' }
];

const chartOptions: echarts.EChartsOption = {
  title: { text: 'Telemetry Signals Rate' },
  xAxis: { type: 'category', data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'] },
  yAxis: { type: 'value' },
  series: [{ data: [120, 200, 150, 80, 70, 110, 130], type: 'line' }]
};

const AppContent: React.FC = () => {
  const { theme, setTheme } = useTheme();
  const [modalOpen, setModalOpen] = useState(false);
  const [toastOpen, setToastOpen] = useState(false);
  const [paletteOpen, setPaletteOpen] = useState(false);
  const [switchChecked, setSwitchChecked] = useState(false);

  const sidebarItems = [
    { label: 'Platform Home', href: '/' },
    { label: 'Vision Analytics', href: '/vision' },
    { label: 'Digital Twin Model', href: '/twin' }
  ];

  const commands = [
    { name: 'Reset Theme to Light Mode', action: () => setTheme('light') },
    { name: 'Set Theme to Dark Mode', action: () => setTheme('dark') },
    { name: 'Trigger Demo Alert Modal', action: () => setModalOpen(true) }
  ];

  const handleFormSubmit = (data: any) => {
    alert(`Form validated successfully! Username: ${data.username}, Email: ${data.email}`);
  };

  const sidebar = (
    <Sidebar items={sidebarItems} currentPath="/" onNavigate={(href) => alert(`Navigation target: ${href}`)} />
  );

  const header = (
    <div className="flex justify-between items-center w-full">
      <span className="text-base font-bold text-text">PRAHARI Control Panel Dashboard</span>
      <div className="flex gap-2.5 items-center">
        <Button size="sm" variant="secondary" onClick={() => setPaletteOpen(true)}>Cmd Menu (Ctrl+K)</Button>
        <select
          value={theme}
          onChange={(e) => setTheme(e.target.value as any)}
          className="text-xs bg-surface border border-border text-text rounded px-2 py-1"
        >
          <option value="light">Light Theme</option>
          <option value="dark">Dark Theme</option>
          <option value="high-contrast">High Contrast</option>
        </select>
      </div>
    </div>
  );

  return (
    <DashboardLayout sidebar={sidebar} header={header}>
      <div className="flex flex-col gap-6">
        {/* KPI Row */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          <KPICard title="Total Camera Inputs" value="48 Cam Streams" change="+4 from last month" trend="up" />
          <KPICard title="Platform Health Status" value="99.8% Online" change="-0.2% latency drift" trend="neutral" />
          <KPICard title="Critical Violations Blocked" value="12 Alerts" change="-3 vs yesterday" trend="down" />
        </div>

        {/* Content grid */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <Card title="Active Incidents Table Summary">
            <DataTable columns={tableColumns} data={tableData} />
          </Card>

          <Card title="ECharts Telemetry Signal Ingestion Chart">
            <EChartWrapper options={chartOptions} />
          </Card>

          <Card title="Dynamic Operator Session Enroller Form">
            <DynamicForm schema={formSchema} fields={formFields} onSubmit={handleFormSubmit} submitLabel="Register Session" />
          </Card>

          <Card title="Alert Action Workspace Actions">
            <div className="flex flex-wrap gap-2.5 items-center">
              <Button onClick={() => setModalOpen(true)}>Open Violation Modal</Button>
              <Button variant="success" onClick={() => setToastOpen(true)}>Show Success Toast</Button>
              <Switch checked={switchChecked} onChange={setSwitchChecked} label="Stream Buffer Telemetries Cache" />
              <Spinner size="sm" />
            </div>
          </Card>
        </div>

        {/* Modals & Overlays */}
        <Modal isOpen={modalOpen} onClose={() => setModalOpen(false)} title="RCA Safety Verification Dialog">
          <p className="text-text/80 mb-4">
            Warning: The perception AI model has detected a restricted boundary intrusion in Reactor Unit Area 2.
            Please verify standard protocol bypass permits before continuing execution.
          </p>
          <div className="flex justify-end gap-2">
            <Button size="sm" variant="secondary" onClick={() => setModalOpen(false)}>Cancel</Button>
            <Button size="sm" variant="danger" onClick={() => { setModalOpen(false); alert('Intrusion flagged.'); }}>Flag Violation</Button>
          </div>
        </Modal>

        {toastOpen && (
          <Toast message="Security logs successfully saved to PostgreSQL datastore." type="success" onClose={() => setToastOpen(false)} />
        )}

        <CommandPalette isOpen={paletteOpen} onClose={() => setPaletteOpen(false)} commands={commands} />
      </div>
    </DashboardLayout>
  );
};

const container = document.getElementById('root');
if (container) {
  const root = createRoot(container);
  root.render(
    <React.StrictMode>
      <ThemeProvider>
        <AppContent />
      </ThemeProvider>
    </React.StrictMode>
  );
}
