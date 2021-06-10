import "../../style/sidebar.css";
import Sticky from "react-sticky-el";
import Suggestions from "./Suggestions";

function Sidebar() {
    return (
        <Sticky topOffset={-20}>
            <div className="sidebar">
                <Suggestions/>
            </div>
        </Sticky>
    );
}

export default Sidebar;