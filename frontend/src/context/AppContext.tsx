// context/AppContext.tsx
import React, { createContext, useContext, useState, useEffect } from 'react';
import type { ReactNode } from 'react';
import type { ModeType, Device, AppContextType } from '../types/app';
import AuthService from '../auth/auth';

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL;

const AppContext = createContext<AppContextType | undefined>(undefined);

export const AppProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [activeMode, setActiveMode] = useState<ModeType>(() => {
    const storedMode = localStorage.getItem('active_mode') as ModeType | null;
    return (storedMode === 'server' || storedMode === 'network') ? storedMode : 'server';
  });

  const [activeDevice, setActiveDevice] = useState<Device | null>(null);
  const [devices, setDevices] = useState<Device[]>([]);
  const [devicesLoading, setDevicesLoading] = useState<boolean>(true);
  const [devicesError, setDevicesError] = useState<string | null>(null);

  // Fetch devices when mode changes
  useEffect(() => {
    const fetchDevices = async () => {
      setDevicesLoading(true);
      setDevicesError(null);
      
      try {
        const endpoint = activeMode === 'server' 
          ? '/api/admin/server/config1/device' 
          : '/api/admin/network/config1/device';
          
        // Use the authenticated request method
        const response = await AuthService.makeAuthenticatedRequest(
          `${BACKEND_URL}${endpoint}`,
          { method: 'GET' }
        );

        if (response.ok) {
          const data = await response.json();
          console.log('Fetched devices:', data);
          
          if (data.status === 'success' && data.devices) {
            setDevices(data.devices);
            
            // Auto-select device if none is selected
            if (data.devices.length > 0 && !activeDevice) {
              const storedDeviceId = localStorage.getItem(`active_device_id_${activeMode}`);
              const deviceToSet = storedDeviceId 
                ? data.devices.find((d: Device) => d.id === storedDeviceId) || data.devices[0]
                : data.devices[0];
              
              setActiveDevice(deviceToSet);
              localStorage.setItem(`active_device_id_${activeMode}`, deviceToSet.id);
            }
          } else {
            setDevices([]);
            setDevicesError('No devices found');
          }
        } else {
          throw new Error(`Failed to fetch devices: ${response.status}`);
        }
      } catch (err) {
        console.error('Error fetching devices:', err);
        setDevices([]);
        setDevicesError(err instanceof Error ? err.message : 'Network error while fetching devices');
      } finally {
        setDevicesLoading(false);
      }
    };

    fetchDevices();
  }, [activeMode, activeDevice]);

  // Reset active device when mode changes
  useEffect(() => {
    setActiveDevice(null);
  }, [activeMode]);

  const updateActiveMode = (mode: ModeType) => {
    console.log('Context: Updating mode to', mode);
    setActiveMode(mode);
    localStorage.setItem('active_mode', mode);
  };

  const updateActiveDevice = (device: Device) => {
    console.log('Context: Updating device to', device.tag);
    setActiveDevice(device);
    localStorage.setItem(`active_device_id_${activeMode}`, device.id);
  };

  const refreshDevices = async () => {
    const endpoint = activeMode === 'server' 
      ? '/api/admin/server/config1/device' 
      : '/api/admin/network/config1/device';
      
    try {
      const response = await AuthService.makeAuthenticatedRequest(
        `${BACKEND_URL}${endpoint}`,
        { method: 'GET' }
      );

      if (response.ok) {
        const data = await response.json();
        if (data.status === 'success' && data.devices) {
          setDevices(data.devices);
          return data.devices;
        }
      }
    } catch (err) {
      console.error('Error refreshing devices:', err);
    }
    return [];
  };

  return (
    <AppContext.Provider value={{
      activeMode,
      updateActiveMode,
      activeDevice,
      updateActiveDevice,
      devices,
      devicesLoading,
      devicesError,
      refreshDevices
    }}>
      {children}
    </AppContext.Provider>
  );
};

// eslint-disable-next-line react-refresh/only-export-components
export const useAppContext = () => {
  const context = useContext(AppContext);
  if (context === undefined) {
    throw new Error('useAppContext must be used within an AppProvider');
  }
  return context;
};
