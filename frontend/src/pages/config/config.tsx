import { useState } from "react";
import Sidebar from "../../components/common/sidebar/sidebar";
import Header from "../../components/common/header/header";
import Config1 from "../../components/server/config/config1/config1";
import Config2 from "../../components/server/config/config2/config2";
import "./config.css";

function Config() {
  const [activeTab, setActiveTab] = useState<"general" | "advanced">("general");

  const handleTabClick = (tab: "general" | "advanced") => {
    setActiveTab(tab);
  };

  return (
    <>
      <Header />
      <div className="container">
        <Sidebar />
        <div className="content">
          <div className="config-content">
            <div className="config-container-1">
              <div className="config-top-bar">
                <div className="config-feature-tabs">
                  <button
                    className={`config-feature-tab ${activeTab === "general" ? "config-feature-tab-active" : ""}`}
                    onClick={() => handleTabClick("general")}
                  >
                    <span className="config-tab-icon">⚙️</span> General Features
                  </button>
                  <button
                    className={`config-feature-tab ${activeTab === "advanced" ? "config-feature-tab-active" : ""}`}
                    onClick={() => handleTabClick("advanced")}
                  >
                    <span className="config-tab-icon">🛡️</span> Advanced
                    Features
                  </button>
                </div>
              </div>
            </div>

            {/* Conditional rendering based on active tab */}
            {activeTab === "general" && <Config1 />}
            {activeTab === "advanced" && <Config2 />}
          </div>
        </div>
      </div>
    </>
  );
}

export default Config;
