import Sidebar from "../../components/common/sidebar/sidebar";
import Resource from "../../components/server/ResourceOptimization/ResourceOptimization";
import Header from "../../components/common/header/header";
import '../../App.css'; 

function Config() {
  return (
    <div className="page-layout">
      <Header />
      <div className="main-section">
        <Sidebar />
        <div className="content-container">
          <Resource />
        </div>
      </div>
    </div>
  );
}

export default Config;
