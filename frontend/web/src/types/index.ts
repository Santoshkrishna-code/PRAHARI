import React from 'react';

export type PageId =
  | 'command-center' | 'ops-intelligence' | 'executive'
  | 'operations' | 'industrial-twin' | 'assets' | 'maintenance'
  | 'permits' | 'incidents' | 'inspections' | 'risk'
  | 'ai-command' | 'agent-orch'
  | 'vision-intel' | 'platform-ops' | 'settings';

export interface NavItem {
  id: PageId;
  label: string;
  icon: React.ReactNode;
  badge?: string;
  accent?: boolean;
}

export interface NavSection {
  label: string;
  items: NavItem[];
}

export interface TelemetryPoint {
  t: string;
  vib: number;
  temp: number;
  psi: number;
  kw: number;
  flow: number;
}

export interface HealthStatus {
  status: string;
  lat: number;
  ts: string;
}

export interface Asset {
  id?: string;
  name: string;
  loc: string;
  type: string;
  health: number;
  rul: string;
  st: string;
  owner: string;
  vib: number | string;
  temp: number | string;
}

export interface Incident {
  id: string;
  title?: string;
  desc: string;
  sev: string;
  asset: string;
  dur?: string;
  st: string;
  time?: string;
}

export interface WorkOrder {
  id: string;
  desc: string;
  asset: string;
  pri: string;
  rul: string;
  st: string;
}

export interface Permit {
  id: string;
  asset: string;
  type: string;
  owner: string;
  st: string;
  expires: string;
}

