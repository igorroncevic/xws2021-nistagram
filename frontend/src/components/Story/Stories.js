import HorizontalScroll from "react-scroll-horizontal";
import Story from "./Story";
import '../../style/Stories.css';

//iz baze iscupaj sve objavljene storije
//Svakkom Story-u proslediti usera koji ga je objavio i media
function Stories() {
    return (
        <div className="stories">
            <HorizontalScroll className="scroll" reverseScroll={false}>
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />
                <Story />

            </HorizontalScroll>
        </div>
    );
}export default Stories;