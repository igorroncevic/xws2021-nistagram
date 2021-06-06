import "../../style/suggestions.css";
import ProfileForSug from "./ProfileForSug";

function Suggestions() {
    return (
        <div className="suggestions">
            <div className="titleContainer">
                <div className="title">Suggestions For You</div>
                <a href="/">See All</a>
            </div>

            <ProfileForSug
                caption="Followed by mapvault + 3 more"
                urlText="Follow"
                iconSize="medium"
                captionSize="small"
                storyBorder={true}
            />
            <ProfileForSug
                caption="Followed by dadatlacak + 1 more"
                urlText="Follow"
                iconSize="medium"
                captionSize="small"
            />
            <ProfileForSug
                caption="Follows you"
                urlText="Follow"
                iconSize="medium"
                captionSize="small"
            />
            <ProfileForSug
                caption="Followed by dadatlacak + 7 more"
                urlText="Follow"
                iconSize="medium"
                captionSize="small"
                storyBorder={true}
            />
            <ProfileForSug
                caption="Followed by mapvault + 4 more"
                urlText="Follow"
                iconSize="medium"
                captionSize="small"
            />
        </div>
    );
}

export default Suggestions;