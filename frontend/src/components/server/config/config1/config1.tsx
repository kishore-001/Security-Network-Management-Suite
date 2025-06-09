import "./config1.css";
import CommandExecution from "./commandexecution";
import SecurityManagement from "./securitymanagement";
import ServerConfig from "./serverconfig";
import ServerOverview from "./serveroverview";
import UserManagement from "./usermanagement";

const Config1 = () => {
  return (
    <>
      <div className="config1-container">
        <h1>Server Configuration Dashboard</h1>
        <div className="grid">
          <CommandExecution />
          <SecurityManagement />
          <ServerConfig />
          <ServerOverview />
          <UserManagement />
        </div>
      </div>
    </>
  );
};

export default Config1;
