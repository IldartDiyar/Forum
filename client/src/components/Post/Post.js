import React, {useState} from 'react';
import './Post.css'
import ThumbUpIcon from '@mui/icons-material/ThumbUp';
import ThumbDownIcon from '@mui/icons-material/ThumbDownAlt';
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import { TextField } from '@mui/material';
import { sendEmotionPost } from '../../api/Posts/Likes';
import { useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { Typography, Alert } from '@mui/material';
import { sendComment } from '../../api/Posts/CreateComment';

function Post(props) {
    const userData = useSelector((state) => state.auth.username);
    const navigate = useNavigate();
    const [comment, setComment] = useState('');
    const [isActive, setIsActive] = useState(true);
    const [likeClicked, setLikeClicked] = useState(props.isLiked);
    const [dislikeClicked, setDislikeClicked] = useState(props.isDisliked);
    const [likes, setLikes] = useState(parseInt(props.countLike));
    const [dislikes, setDislikes] = useState(parseInt(props.countDislike));
    const [commentError, setCommentError] = useState("");
    const [emotionError, setEmotionError] = useState(null);
    const handleNewComment = () => { //Toggling Add comment button
      setIsActive(current => !current)
    }

    const handleCommentStatus = () => {
      if (commentError){
          return (  
              <Alert severity="error">
                {commentError.message || commentError.toString()}
              </Alert>
          )
      }
    }
    const handleEmotionStatus = () => {
      if (emotionError){
          return (  
              <Alert severity="error">
                {emotionError.message || emotionError.toString()}      
              </Alert>
          )
      }
    }
    
      function handleLikeClick() { //Like logic
        if (userData === null) {
          return; // do nothing if user is not authorized
        }
        if (!likeClicked) {
          sendEmotionPost(1, parseInt(props.query), setEmotionError, navigate);
          setLikes(prevCount => prevCount + 1);
          setLikeClicked(true);
          if (dislikeClicked) {
            setDislikes(prevCount => prevCount - 1);  
            setDislikeClicked(false);
          }
        } else{
          sendEmotionPost(-1, parseInt(props.query), setEmotionError, navigate);
          setLikes(prevCount => prevCount - 1);
          setLikeClicked(false);
        }
      }
    
      function handleDislikeClick() { //Dislike logic
        if (userData === null) {
          return; // do nothing if user is not authorized
        }
        if (!dislikeClicked) {
          sendEmotionPost(0, parseInt(props.query), setEmotionError, navigate);
          setDislikes(prevCount => prevCount + 1);
          setDislikeClicked(true);
          if (likeClicked) {
            setLikes(prevCount => prevCount - 1);
          
            setLikeClicked(false);
          }
        } else{
          sendEmotionPost(-1, parseInt(props.query), setEmotionError, navigate);
          setDislikes(prevCount => prevCount - 1);
          setDislikeClicked(false);
        }
      }

    return (
      <>
      <div className='post-container'>
        {/* Header */}
        <div className='post-header'>
          <div className='post-header-date'>
            <p>{props.date}</p>
          </div>
          <div className='post-header-title'>
            <p>{props.title}</p>
          </div>
          <div className='post-header-author'>
            <p>Author: {props.author}</p>
          </div>
        </div>
        
        {/* Body */}
        <div className='post-body'>
            <p>{props.content}</p>
        </div>
        
        {/* Footer */}
        <div className='post-footer'>
          <div className='post-footer-comments'>
            Some data for the future
          </div> 
        
          
          
          {userData!==null && 
            <div className='post-footer-add-comment'>
                <AddCircleOutlineIcon
                  fontSize='large'
                  onClick={handleNewComment}
                  className='add-comment-button'
                />
            </div>
          }  
          
          <div className='post-footer-likes'>
            <div className='post-footer-likes-icons'>
            <ThumbUpIcon 
              onClick={userData!==null ? handleLikeClick: undefined} 
              sx={{
                ...(userData!==null && {'&:hover':{
                  cursor: 'pointer',
                  color: '#58da47'
                  },
                  color: likeClicked ? 'white' : undefined
                })
            }}
            />
            <p>{likes}</p>
            
            <ThumbDownIcon
              onClick={userData!==null? handleDislikeClick: undefined}
              sx={{
                ...(userData!==null && {'&:hover':{
                  cursor: 'pointer',
                  color: '#e75217'
                  },
                  color: dislikeClicked ? 'white' : undefined
                })
              }}
            />
            <p>{dislikes}</p>
            </div>
          </div>
        </div>
        
        {/* Input Form for a new comment */}
      </div>
      { userData!==null &&
        <div className='add-comment-input' style={{display: isActive ? "block" : "none"}}>
          <form className='add-comment-form'>  
            <TextField 
              id="filled-basic"
              label="Add Comment"
              sx={{
                backgroundColor:"white",
                borderRadius: "10px",
                width: "90vw",
                height: "100%"
              }}
              onChange={(e) => setComment(e.target.value)}
            />
            <button className="custom-btn btn-15" onClick={(e) => sendComment(e, userData, comment, parseInt(props.query), navigate, setCommentError)}>Add</button>
          </form>
        </div>
    }
            <Typography component="h1" variant="h5" sx={{color:'black'}}>
                {handleCommentStatus()}
                {/* CREATE COMMENT ERROR HANDLING */}
            </Typography>
            <Typography component="h1" variant="h5" sx={{color:'black'}}>
                {handleEmotionStatus()} 
                {/* LIKE ERROR HANDLING */}
            </Typography>
    </>
    );
  }
export default Post;