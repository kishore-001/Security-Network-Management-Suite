
import Sidebar from "../../components/common/sidebar/sidebar";
import Header from "../../components/common/header/header";
import Config2 from "../../components/server/config/config1/config1";
import '../../App.css'; // Assuming App.css is at src/App.css

function Config() {
  return (
    <div className="page-layout">
      <Header />
      <div className="main-section">
        <Sidebar />
        <div className="content-container">
          <Config2 />
        </div>
      </div>
    </div>
  );
}

export default Config;
