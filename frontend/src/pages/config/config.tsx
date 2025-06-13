import Sidebar from "../../components/common/sidebar/sidebar";
import Header from "../../components/common/header/header";
import Config1 from "../../components/server/config/config1/config1";
import "./config.css"; // Assuming App.css is at src/App.css

function Config() {
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
                  <button className="config-feature-tab">
                    <span className="tab-icon"></span> General Features
                  </button>
                  <button className="config-feature-tab">
                    <span className="config-tab-icon">🛡️</span> Advanced
                    Features
                  </button>
                </div>
              </div>
            </div>
            <Config1 />
          </div>
        </div>
      </div>
    </>
  );
}

export default Config;
