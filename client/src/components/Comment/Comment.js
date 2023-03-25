import ThumbUpIcon from '@mui/icons-material/ThumbUp';
import ThumbDownIcon from '@mui/icons-material/ThumbDownAlt';
import { useState } from 'react';
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { Alert } from '@mui/material';
import { sendEmotionComment } from '../../api/Posts/Likes';
import { useSelector } from 'react-redux';
import { useNavigate } from 'react-router';

function CommentComponent(props) {
  const [likeClicked, setLikeClicked] = useState(props.isLiked);
  const [dislikeClicked, setDislikeClicked] = useState(props.isDisliked);
  const [likes, setLikes] = useState(parseInt(props.countLike));
  const [dislikes, setDislikes] = useState(parseInt(props.countDislike));
  const [emotionError, setEmotionError] = useState(null);
  const userData = useSelector((state) => state.auth.username);
  const navigate = useNavigate();
      function handleLikeClick() { // Like Comment Logic
        if (!likeClicked) {
          sendEmotionComment(1, parseInt(props.commentID), setEmotionError, navigate);
          setLikes(prevCount => prevCount + 1);
          setLikeClicked(true);
          if (dislikeClicked) {
            setDislikes(prevCount => prevCount - 1);  
            setDislikeClicked(false);
          }
        } else{
          sendEmotionComment(-1, parseInt(props.commentID), setEmotionError, navigate);
          setLikes(prevCount => prevCount - 1);
          setLikeClicked(false);
        }
      }
    
      function handleDislikeClick() { // Dislike Comment Logic
        if (!dislikeClicked) {
          sendEmotionComment(0, parseInt(props.commentID), setEmotionError, navigate);
          setDislikes(prevCount => prevCount + 1);
          setDislikeClicked(true);
          if (likeClicked) {
            setLikes(prevCount => prevCount - 1);
          
            setLikeClicked(false);
          }
        } else{
          sendEmotionComment(-1, parseInt(props.commentID), setEmotionError, navigate);
          setDislikes(prevCount => prevCount - 1);
          setDislikeClicked(false);
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
      return (
        
        <Box
          sx={{
            backgroundColor: 'rgb(84, 107, 156)',
            position: 'relative',
            m: 1,
            height: '15vh',
            borderRadius: 1,
            display: 'flex',
            flexDirection: 'column',
          }}
        >
            
          {/* Header */}
          <Box
            sx={{
              minHeight: '20%',
              fontFamily: 'Bebas Neue',
              display: 'flex',
              color: 'black',
            }}
          >
            <Box sx={{ height: '100%', width: '100%', display: 'flex', alignItems: 'center' }}>
              <Typography sx={{ ml: 2, fontWeight: 600 }}>{props.date}</Typography>
            </Box>
            <Box
              sx={{
                height: '100%',
                width: '100%',
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
              }}
            >
            </Box>
            <Box
              sx={{
                height: '100%',
                width: '100%',
                display: 'flex',
                justifyContent: 'flex-end',
                alignItems: 'center',
              }}
            >
              <Typography sx={{ mr: 1, fontWeight: 600 }}>Author: {props.author}</Typography>
            </Box>
          </Box>
      
          {/* Body */}
          <Box
            sx={{
              backgroundColor: 'rgb(12, 17, 26)',
              minHeight: '60%',
              overflow: 'auto',
              letterSpacing: 1,
              color: 'rgb(193, 255, 255)',
            }}
          >
            <Typography sx={{ p: 2, maxWidth: '95%', maxHeight: '100%' }}>{props.content}</Typography>
          </Box>
      
          {/* Footer */}
          <Box sx={{ minHeight: '20%', fontFamily: 'Bebas Neue', display: 'flex', alignItems: 'center' }}>
            <Box sx={{ height: '100%', width: '100%', display: 'flex', alignItems: 'center', pl: 2 }}>
              <Box>
                    {handleEmotionStatus()} 
                    {/* ERROR HANDLER */}
              </Box>
            </Box>
      
            <Box
              sx={{
                height: '100%',
                width: '100%',
                color: 'rgb(193, 255, 255)',
                display: 'flex',
                alignItems: 'center',
              }}
            >
            <Box sx={{ display: 'inline', px: 0.5 }}>
            
          </Box>
          <Box sx={{ display: 'flex', alignItems: 'center', mr: 2, ml: 'auto' }}>
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
            <Box sx={{ display: 'inline', px: 0.5 }}>
              <Typography>{likes}</Typography>
            </Box>
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
            <Box sx={{ display: 'inline', px: 0.5 }}>
              <Typography>{dislikes}</Typography>
            </Box>
        
          </Box>
        </Box>
      </Box>
       

      </Box>
      );
      
  }
export default CommentComponent;

