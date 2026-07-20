import React from 'react';
import { usePermissions } from '@prahari/utils/permissions.tsx';

interface RoleGuardProps {
  requiredRole: string;
  fallback?: React.ReactNode;
  children: React.ReactNode;
}

export const RoleGuard: React.FC<RoleGuardProps> = ({
  requiredRole,
  fallback = <div className="p-6 text-center text-danger text-sm font-semibold">Bypassed session: Access Denied.</div>,
  children
}) => {
  const { hasPermission } = usePermissions();

  if (!hasPermission(requiredRole)) {
    return <>{fallback}</>;
  }

  return <>{children}</>;
};
