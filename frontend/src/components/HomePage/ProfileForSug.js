import "../../style/profileSug.css";
import ProfileIcon from "../ProfileComponent/ProfileIcon";

function ProfileForSug(props) {
    const { username, caption,  urlText, iconSize,captionSize, storyBorder,hideAccountName,image,} = props;

    return (
        <div className="profile">
            <ProfileIcon iconSize={iconSize} storyBorder={storyBorder} image={image} />
            {('maja' || caption) && !hideAccountName && (
                <div className="textContainer">
                    <span className="accountName">Maja</span>
                    <span className={`caption ${captionSize}`}>{caption}</span>
                </div>
            )}
            <a href="/">{urlText}</a>
        </div>
    );
}

export default ProfileForSug;