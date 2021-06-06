import "../../style/suggestions.css";
import ProfileForSug from "./ProfileForSug";

function Suggestions() {
    return (
        <div className="suggestions">
            <div className="titleContainer">
                <div className="title">Suggestions For You</div>
                <a href="/">See All</a>
            </div>

            <ProfileForSug caption="Health" urlText="Follow" iconSize="medium" captionSize="small"  storyBorder={true} />
            <ProfileForSug caption="Sport" urlText="Follow"  iconSize="medium" captionSize="small" />
            <ProfileForSug caption="Follows you" urlText="Follow"  iconSize="medium" captionSize="small" />
            <ProfileForSug caption="Baby"  urlText="Follow" iconSize="medium"  captionSize="small" storyBorder={true} />
            <ProfileForSug caption="Sport" urlText="Follow" iconSize="medium"  captionSize="small" />
        </div>
    );
}

export default Suggestions;