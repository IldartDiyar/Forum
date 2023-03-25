import { baseURL } from "../API";


// IMAGE UPLOAD
export const uploadImage = async (e, image, setStatus, setError, navigate) => {
    e.preventDefault();
    if (image !== null){
        const formData = new FormData();
        formData.append('image', image);
        
        await fetch(`${baseURL}/auth/upload_avatar`, {
        method: 'POST',
        credentials: 'include',
        body: formData,
        }).then(async (response) => {
          if (!response.ok){
            if (response.status === 401){
              setStatus('');
              setError('');
              const error = new Error('You are not authorized for image upload');
              error.status = response.status;
              navigate("/signin", {state: {error}})
           } else if (response.status === 400){
              setStatus('');  
              const responseBody = await response.text();
              if (responseBody){
                setError(responseBody);  
              } else {
                setError("Error uploading an image")
              }
           } else {
              setStatus('');
              setError(`Couldn't upload an image. Status: ${response.statusText}`);
            }
          } else if (response.ok) {
            setStatus('Image uploaded');
            setError('');
          }
        }) 
    }
}
  // POTENTIAL FUTURE FETCHES FOR USER INFORMATION EDITING.