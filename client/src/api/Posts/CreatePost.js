//CREATE POST
import { baseURL } from "../API";

export const createPost = async(e, text, category, title, userData, setStatus, navigate) => {
    e.preventDefault();
        
    await fetch(`${baseURL}/post/create_post`, {
        method: 'POST',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Content: text,
          Category: category,
          Title: title,
          Author: userData,
          ImageData: null //FOR THE FUTURE
        }),
      }).then(async (response) => {
        if (!response.ok) {
          if (response.status === 400){
            const responseBody = await response.text();
            setStatus(responseBody)
          } else if (response.status === 401){
            setStatus('');
            const error = new Error('Only authorized users are allowed to create posts');
            error.status = response.status;
            navigate("/signin", {state: {error}})
          } else {
            setStatus(`Couldn't create a post. Status: ${response.statusText}`)
          }
        } else {
          setStatus('');
          return navigate("/");  
        }   
      })
      
}