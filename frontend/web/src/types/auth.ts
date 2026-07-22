export type UserRole =
  | 'Administrator'
  | 'Safety Officer'
  | 'Plant Manager'
  | 'Reliability Engineer'
  | 'Maintenance Engineer'
  | 'Operator'
  | 'Auditor'
  | 'Executive';

export interface UserSession {
  id: string;
  name: string;
  email: string;
  role: UserRole;
  orgName: string;
  plantName: string;
  avatar: string;
  token: string;
}

export interface OrganizationSetup {
  companyName: string;
  industry: string;
  country: string;
  timezone: string;
  plantName: string;
  location: string;
  assetSource: string;
  teamEmail: string;
  aiModel: string;
}
