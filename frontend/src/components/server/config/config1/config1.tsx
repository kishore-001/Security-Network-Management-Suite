import React from "react";
import "./config1.css";
import ServerConfiguration from "./serverconfiguration";
import CommandExecution from "./commandexecution";
import UserManagement from "./usermanagement";
import SecurityManagement from "./securitymanagement";
import ServerOverview from "./serveroverview";

const Config1: React.FC = () => {
  return (
    <div className="config1-container-2">
      {/* Title & Description */}
      <div className="config1-header">
        <div className="config-icon" />
        <h1>Server Configuration Management</h1>
        <p>
          Comprehensive server administration, user management, and security
          configuration in one unified platform
        </p>
      </div>

      {/* Four Horizontal Cards */}
      <div className="card-row-horizontal">
        <ServerConfiguration />
        <CommandExecution />
        <UserManagement />
        <SecurityManagement />
      </div>

      {/* Overview */}
      <div className="overview-wrapper">
        <ServerOverview />
      </div>
    </div>
  );
};

export default Config1;
