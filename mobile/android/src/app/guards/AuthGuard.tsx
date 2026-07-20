import React from 'react';
import { useAuth } from '@prahari/utils/auth.tsx';

interface AuthGuardProps {
  fallback: React.ReactNode;
  children: React.ReactNode;
}

export const AuthGuard: React.FC<AuthGuardProps> = ({ fallback, children }) => {
  const { isAuthenticated } = useAuth();

  if (!isAuthenticated) {
    return <>{fallback}</>;
  }

  return <>{children}</>;
};
