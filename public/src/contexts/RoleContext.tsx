import React, { createContext, useContext, useEffect, useState, ReactNode } from 'react';
import { roles } from '../constants/roles';
import { verifyToken } from '../api/authClient';

export interface RoleContextType {
  role: string; // ロールを表す文字列型
  loginAsUser: () => void;
  loginAsAdmin: () => void;
  logout: () => void;
  setRole: (role: string) => void;
}

const defaultValue: RoleContextType = {
  role: roles.GUEST,
  loginAsUser: () => {},
  loginAsAdmin: () => {},
  logout: () => {},
  setRole: () => {},
};

export const RoleContext = createContext<RoleContextType>(defaultValue);

export const RoleProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [role, setRole] = useState(roles.GUEST);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const loginAsUser = () => setRole(roles.USER);
  const loginAsAdmin = () => setRole(roles.ADMIN);
  const logout = () => setRole(roles.GUEST);

  const refreshRole = async () => {
    try {
      setLoading(true);
      const accessToken = localStorage.getItem('access_token');
      if (!accessToken) throw new Error('No access token found');

      const { role } = await verifyToken(accessToken);
      console.log(role)
      setRole(role);
      setError(null);
    } catch (err: any) {
      console.error(err)
      setError(err.message || 'Failed to verify token');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    refreshRole(); // 初回ロード時にロールを取得
  }, []);

  return (
    <RoleContext.Provider value={{ role, loginAsUser, loginAsAdmin, logout, setRole }}>
      {children}
    </RoleContext.Provider>
  );
};

export const useRole = () => useContext(RoleContext);
