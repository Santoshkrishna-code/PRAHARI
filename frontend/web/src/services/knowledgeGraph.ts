export interface GraphEntity {
  id: string;
  name: string;
  type: 'Organization' | 'Plant' | 'Area' | 'Line' | 'Asset' | 'Sensor' | 'Camera' | 'Maintenance' | 'Incident' | 'Permit' | 'Report';
  status: 'Healthy' | 'Warning' | 'Critical' | 'Active';
  details: string;
  relationships: { entityId: string; relation: string }[];
}

class KnowledgeGraphService {
  private entities: Map<string, GraphEntity> = new Map();

  constructor() {
    this.seedGraph();
  }

  private seedGraph() {
    const data: GraphEntity[] = [
      {
        id: 'org-alpha',
        name: 'Alpha Chemical Refinery Inc.',
        type: 'Organization',
        status: 'Healthy',
        details: 'Enterprise Tenant — Oil & Gas Processing',
        relationships: [{ entityId: 'plant-gulf', relation: 'operates' }],
      },
      {
        id: 'plant-gulf',
        name: 'Plant Alpha (Gulf Coast)',
        type: 'Plant',
        status: 'Healthy',
        details: 'Primary Chemical Refining Complex — Houston, TX',
        relationships: [
          { entityId: 'org-alpha', relation: 'belongs_to' },
          { entityId: 'area-rx-b', relation: 'contains' },
        ],
      },
      {
        id: 'area-rx-b',
        name: 'Reactor Complex B',
        type: 'Area',
        status: 'Warning',
        details: 'High-Pressure Recirculation Complex',
        relationships: [
          { entityId: 'plant-gulf', relation: 'belongs_to' },
          { entityId: 'line-dc101', relation: 'contains' },
        ],
      },
      {
        id: 'line-dc101',
        name: 'Recirculation Loop DC-101',
        type: 'Line',
        status: 'Warning',
        details: 'Slurry Recirculation Pipeline',
        relationships: [
          { entityId: 'area-rx-b', relation: 'belongs_to' },
          { entityId: 'PUMP-P102', relation: 'houses' },
        ],
      },
      {
        id: 'PUMP-P102',
        name: 'Pump P-102 (Slurry Recirculation)',
        type: 'Asset',
        status: 'Warning',
        details: 'Centrifugal Slurry Pump — RUL: 18 days',
        relationships: [
          { entityId: 'line-dc101', relation: 'located_in' },
          { entityId: 'sensor-vp102', relation: 'monitored_by' },
          { entityId: 'cam-002', relation: 'scanned_by' },
          { entityId: 'wo-7821', relation: 'has_work_order' },
          { entityId: 'inc-0447', relation: 'associated_incident' },
          { entityId: 'ptw-8902', relation: 'has_active_permit' },
        ],
      },
      {
        id: 'sensor-vp102',
        name: 'Vibration Probe VP-102',
        type: 'Sensor',
        status: 'Warning',
        details: 'ISO 10816 Velocity Probe — 11.8 mm/s',
        relationships: [{ entityId: 'PUMP-P102', relation: 'monitors' }],
      },
      {
        id: 'cam-002',
        name: 'Camera CAM-002 (Jetson AGX-04)',
        type: 'Camera',
        status: 'Warning',
        details: 'YOLOv8 Edge Camera — Restricted Zone B',
        relationships: [{ entityId: 'PUMP-P102', relation: 'scans' }],
      },
      {
        id: 'wo-7821',
        name: 'Work Order WO-7821',
        type: 'Maintenance',
        status: 'Warning',
        details: 'Bearing Replacement & Lubrication Service (Overdue)',
        relationships: [{ entityId: 'PUMP-P102', relation: 'targets' }],
      },
      {
        id: 'inc-0447',
        name: 'Incident INC-2026-0447',
        type: 'Incident',
        status: 'Warning',
        details: 'Vibration Excursion & PPE Non-Compliance',
        relationships: [{ entityId: 'PUMP-P102', relation: 'investigates' }],
      },
      {
        id: 'ptw-8902',
        name: 'Safe Work Permit PTW-8902',
        type: 'Permit',
        status: 'Active',
        details: 'Hot Work Permit — Isolation Lock Valve V-88',
        relationships: [{ entityId: 'PUMP-P102', relation: 'isolates' }],
      },
    ];

    data.forEach(e => this.entities.set(e.id, e));
  }

  public getEntity(id: string): GraphEntity | undefined {
    return this.entities.get(id);
  }

  public getRelated(id: string): GraphEntity[] {
    const entity = this.entities.get(id);
    if (!entity) return [];

    return entity.relationships
      .map(r => this.entities.get(r.entityId))
      .filter((e): e is GraphEntity => e !== undefined);
  }

  public getAll(): GraphEntity[] {
    return Array.from(this.entities.values());
  }
}

export const knowledgeGraph = new KnowledgeGraphService();
