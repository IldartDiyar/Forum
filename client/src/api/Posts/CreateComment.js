import { baseURL } from "../API";

export const sendComment = async (event, userData, comment, postID, navigate, setCommentError) => {
    event.preventDefault();
        await fetch(`${baseURL}/post/create_comment`, {
          method: 'POST',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            author: userData,
            body: comment,
            postId: postID
          }),
        }).then(async(response) => {
          
          if (!response.ok) {
            if (response.status === 400){
              const responseBody = await response.text();
              setCommentError(responseBody)
            } else if (response.status === 401){
              setCommentError("");
              const error = new Error('Only authorized users are allowed to send comments');
              error.status = response.status;
              navigate("/signin", {state: {error}})
            } else {
              setCommentError(`Couldn't create a comment. Status: ${response.statusText}`)
            }
          } else {
            //Navigate to the same post page
            setCommentError("");
            navigate(0);
          }
        })

}