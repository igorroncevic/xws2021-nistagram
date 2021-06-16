import "./../../style/PostHeader.css";
import ProfileIcon from "../ProfileComponent/ProfileIcon";

const PostHeader = (props) => {
    const {
      username,
      caption,
      captionSize,
      iconSize,
      hideUsername,
      image,
    } = props;
    
    return (
      <div className="post-header">
        <ProfileIcon
          iconSize={iconSize}
          //storyBorder={storyBorder}
          image={image}
        />
        {(username || caption) && !hideUsername && (
          <div className="textContainer">
            {/* Make username clickable */}
            <span className="username">{username}</span>
            <span className={`caption ${captionSize}`}>{caption}</span>
          </div>
        )}
      </div>
    );
  }
  
  export default PostHeader;