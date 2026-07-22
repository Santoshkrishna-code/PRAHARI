import React from 'react';
import {
  LayoutDashboard, BarChart3, FileText, Radio, Box, Package, Wrench,
  Shield, AlertTriangle, ClipboardCheck, Target, Brain, Cpu, Eye, Server, Settings
} from 'lucide-react';
import { NavSection } from '../../types';

export const NAV_CONFIG: NavSection[] = [
  {
    label: 'Executive',
    items: [
      { id: 'command-center', label: 'Command Center', icon: <LayoutDashboard size={17} /> },
      { id: 'ops-intelligence', label: 'Operations Intelligence', icon: <BarChart3 size={17} /> },
      { id: 'executive', label: 'Executive Insights', icon: <FileText size={17} /> },
    ]
  },
  {
    label: 'Operations',
    items: [
      { id: 'operations', label: 'Operations Center', icon: <Radio size={17} /> },
      { id: 'industrial-twin', label: 'Industrial Twin', icon: <Box size={17} /> },
      { id: 'assets', label: 'Assets', icon: <Package size={17} /> },
      { id: 'maintenance', label: 'Maintenance', icon: <Wrench size={17} />, badge: '3' },
    ]
  },
  {
    label: 'Safety',
    items: [
      { id: 'permits', label: 'Safe Work Permits', icon: <Shield size={17} /> },
      { id: 'incidents', label: 'Incidents', icon: <AlertTriangle size={17} />, badge: '1' },
      { id: 'inspections', label: 'Inspections', icon: <ClipboardCheck size={17} /> },
      { id: 'risk', label: 'Risk Assessment', icon: <Target size={17} /> },
    ]
  },
  {
    label: 'AI',
    items: [
      { id: 'ai-command', label: 'AI Command Center', icon: <Brain size={17} />, accent: true },
      { id: 'agent-orch', label: 'Agent Orchestration', icon: <Cpu size={17} /> },
    ]
  },
  {
    label: 'Platform',
    items: [
      { id: 'vision-intel', label: 'Vision Intelligence', icon: <Eye size={17} /> },
      { id: 'platform-ops', label: 'Platform Operations', icon: <Server size={17} /> },
      { id: 'settings', label: 'Settings', icon: <Settings size={17} /> },
    ]
  },
];
