import HorizontalScroll from "react-scroll-horizontal";
import Story from "./Story";

function Stories() {
    const stories={
        display: 'flex',
        alignContent: 'center',
        backgroundColor: '#ffffff',
        border: '1px solid #dbdbdb',
        borderRadius: '3px',
        margin: '2em 0 1.5em 0',
        height: '7.4em',
        overflow: 'hidden',
    }
    return (
        <div style={stories}>
            <HorizontalScroll className="scroll" reverseScroll={true}>
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