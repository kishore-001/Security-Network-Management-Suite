import React, { useState } from 'react';
import Header from '../../components/common/header/header';
import Sidebar from '../../components/common/sidebar/sidebar';
import SettingsTabs from './SettingsTabs';
import Health from '../../components/common/health/health';
import UserManagement from './UserManagement/UserManagement';
import AdditionalConfig from './AdditionalConfig/AdditionalConfig';
import './settings.css';

const SettingsPage: React.FC = () => {
  const [activeTab, setActiveTab] = useState<'user' | 'config'>('user');

  return (
  <div className="settings-layout">
  <Header />

  <div className="main-content">
    <div className="left-panel">
         <Health />
     <div className="sidebar">
      <Sidebar />
     </div>
      
    </div>

    <div className="settings-container">
      <SettingsTabs activeTab={activeTab} setActiveTab={setActiveTab} />
      {activeTab === 'user' ? <UserManagement /> : <AdditionalConfig />}
    </div>
  </div>
</div>

  );
};

export default SettingsPage;
