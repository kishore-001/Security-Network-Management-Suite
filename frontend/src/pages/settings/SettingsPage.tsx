import React, { useState } from 'react';
import SettingsTabs from './SettingsTabs';
import UserManagement from './UserManagement/UserManagement';
import AdditionalConfig from './AdditionalConfig/AdditionalConfig';
import './settings.css';

const SettingsPage: React.FC = () => {
  const [activeTab, setActiveTab] = useState<'user' | 'config'>('user');

  return (
    <div className="settings-container">
      <SettingsTabs activeTab={activeTab} setActiveTab={setActiveTab} />
      {activeTab === 'user' ? <UserManagement /> : <AdditionalConfig />}
    </div>
  );
};

export default SettingsPage;