import { baseURL } from "../API";

export const sendEmotionComment = async (emotion, commentID, setEmotionError, navigate) => {
    try{
      const response = await fetch(`${baseURL}/emotian/comment`, {
        method: 'POST',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          CommentID: commentID,
          Islike: emotion
        }),
      });
      
      if (!response.ok) {
        if (response.status === 401) {
          const error = new Error('Only authorized users are allowed to give reactions');
          error.status = response.status;
          navigate("/signin", {state: {error}})
        } else {
          const error = new Error(`Error while liking the post. Status: ${response.statusText}`);
          error.status = response.status;
          throw error;
        }
      } else {
        setEmotionError(null);
      }
    }catch (err) {
      console.log(err);
      setEmotionError(err);
    }
  }
  
  export const sendEmotionPost = async (emotion, postID, setEmotionError, navigate) => {
    try{
          const response = await fetch(`${baseURL}/emotian/post`, {
            method: 'POST',
            credentials: 'include',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
              PostID: postID,
              Islike: emotion
            }),
          });
          if (!response.ok) {  
            if (response.status === 401) {
              const error = new Error('Only authorized users are allowed to give reactions');
              error.status = response.status;
              navigate("/signin", {state: {error}})
            } else {
              const error = new Error(`Error while liking the post. Status: ${response.statusText}`);
              error.status = response.status;
              throw error;
            }
          } else {
            setEmotionError(null);
          }
        }catch (err) {
          console.log(err);
          setEmotionError(err);
        }
    }