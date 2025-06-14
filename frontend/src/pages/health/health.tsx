import Sidebar from "../../components/common/sidebar/sidebar";
import Header from "../../components/common/header/header";
import "../../App.css";
//import Health from "../../components/server/health/health";
function Config() {
  return (
    <div className="page-layout">
      <Header />
      <div className="main-section">
        <Sidebar />
        <div className="content-container"></div>
      </div>
    </div>
  );
}

export default Config;
