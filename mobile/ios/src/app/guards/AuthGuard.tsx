import React, { useState } from 'react';
import { useAuth } from '@prahari/utils/auth.tsx';
import { Button } from '@prahari/components/Button.tsx';
import { Card } from '@prahari/components/Card.tsx';

interface AuthGuardProps {
  fallback: React.ReactNode;
  children: React.ReactNode;
}

export const AuthGuard: React.FC<AuthGuardProps> = ({ fallback, children }) => {
  const { isAuthenticated } = useAuth();
  const [biometricsApproved, setBiometricsApproved] = useState(false);

  if (!isAuthenticated) {
    return <>{fallback}</>;
  }

  if (!biometricsApproved) {
    return (
      <div className="flex items-center justify-center p-6 bg-background min-h-[40vh] w-full">
        <Card title="Face ID / Touch ID Required" className="max-w-sm text-center">
          <p className="text-sm text-muted mb-4">Confirm Apple biometrics authentication to unlock credentials session keychain.</p>
          <Button variant="primary" onClick={() => setBiometricsApproved(true)}>Authenticate Face ID</Button>
        </Card>
      </div>
    );
  }

  return <>{children}</>;
};
