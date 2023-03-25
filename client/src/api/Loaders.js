import { loginSuccess, logout } from '../features/auth'
import {store} from '../store'
import { redirect } from 'react-router-dom';
import { baseURL } from './API';
import { signOutOnError } from './Authorization/Authorization';

// Token Check at page refresh or first visit
export const checkAuth =( async () => {
    try{
      const response = await fetch(`${baseURL}/auth/check_token`, {
        headers: {
          'Accept': 'application/json',
          'Credentials': 'include'
        },
        method: "GET",
        credentials: 'include',
      })
      if (!response.ok){
        if (response.status === 401) {
          store.dispatch(logout());
          return null;
        } else{
          const error = new Error(response.statusText);
          error.status = response.status;
          throw error;
        }
      } else if (response.ok){
        const data = await response.json(); 
        if (data.Username){ // Check if I get username in my response
          store.dispatch(loginSuccess({ username: data.Username }));
        } else { // Handle if there is no username in the response
          store.dispatch(logout());
          signOutOnError();
          const error = new Error("Error while getting userdata from the server"); // Very unlikely to happen
          error.status = 500;
          throw error;
        }
        return null;
      }
    } catch(error){
      console.log(error);
      throw error;
    }
  })  //OK
//PROFILE PAGE LOADER
export const profileData =( async () => {
  try{
    const response = await fetch(`${baseURL}/personal_info`, {
      headers: {
        'Accept': 'application/json',
        'Credentials': 'include'
      },
      method: "GET",
      credentials: 'include',
    })
    if (!response.ok){
      if (response.status === 401){
        store.dispatch(logout());
        return redirect("/signin")
      } else{
        const error = new Error(response.statusText);
        error.status = response.status;
        throw error;
      }
    } else if (response.ok){
      const data = await response.json();
      return {data};
    }
  } catch(error){
    console.log(error);
    throw error;
  }
})  //OK

// LOADING POST DATA AND ITS COMMENTS
export const postLoader = async ({request}) => {
  try {
    const url = new URL(request.url);
    const searchParams = url.searchParams;
    const id = searchParams.get('id');
      
    const [postResponse, commentsResponse] = await Promise.all([
      fetch(`${baseURL}/post/get_post/?id=${id}`, {
        headers: {
          Accept: "application/json",
          Credentials: "include",
        },
        method: "GET",
        credentials: "include",
      }),
      fetch(`${baseURL}/post/get_comment/?id=${id}`, {
        headers: {
          Accept: "application/json",
          Credentials: "include",
        },
        method: "GET",
        credentials: "include",
      })
    ]);
      
    if (!postResponse.ok) {
      const error = new Error(`Could not fetch the post. Status: ${postResponse.statusText}`);
      error.status = postResponse.status;
      throw error;
    }
    
    if (!commentsResponse.ok) {
      const error = new Error(`Could not fetch the comments. Status: ${commentsResponse.statusText}`);
      error.status = commentsResponse.status;
      throw error;
    }
    
    const post = await postResponse.json();
    const comments = await commentsResponse.json();
    return { post, comments };
  } catch (error) {
    console.log(error);
    throw error;
  }
} //OK

  //Create-Post
  export const createPostLoader = async () => {
    try{
      const [checkResponse, categoriesResponse] = await Promise.all([
        fetch(`${baseURL}/auth/check_token`, {
          headers: {
            'Accept': 'application/json',
            'Credentials': 'include'
          },
          method: "GET",
          credentials: 'include',
        }),
        fetch(`${baseURL}/post/get_post_category`, {
          headers: {
            Accept: "application/json",
            Credentials: "include",
          },
          method: "GET",
          credentials: "include",
        })
      ]);
        
        if (!checkResponse.ok) {
          store.dispatch(logout());
          if (checkResponse.status === 401){
            return redirect("/signin");  
          } else {
            const error = new Error(`Status: ${checkResponse.statusText}`)
            error.status = checkResponse.status;
            throw error;
          }
        } else if (checkResponse.ok){
          const data = await checkResponse.json(); 
          store.dispatch(loginSuccess({ username: data.Username }));
        }
        if (!categoriesResponse.ok) {
          const error = new Error(`Could not fetch the categories. Status: ${categoriesResponse.statusText}`);
          error.status = categoriesResponse.status;
          throw error;
        }
        const categories = await categoriesResponse.json();
        return {categories};  
    } catch(error){
      console.log(error);
      throw error;
    }
  } //OK

  
  //Loading all posts and categories for home page, myposts and liked posts
  const fetchPostsAndCategories = async (postsEndpoint, currentQuery) => {
    try {
      const [postsResponse, categoriesResponse] = await Promise.all([
        fetch(`${baseURL}${postsEndpoint}${currentQuery}`, {
          headers: {
            Accept: "application/json",
            Credentials: "include",
          },
          method: "GET",
          credentials: "include",
        }),
        fetch(`${baseURL}/post/get_post_category`, {
          headers: {
            Accept: "application/json",
            Credentials: "include",
          },
          method: "GET",
          credentials: "include",
        }),
      ]);
  
      if (!postsResponse.ok) {
        const error = new Error(`Could not fetch posts. Status: ${postsResponse.statusText}`);
        error.status = postsResponse.status;
        throw error;
      }
  
      if (!categoriesResponse.ok) {
        const error = new Error(`Could not fetch the categories. Status: ${categoriesResponse.statusText}`);
        error.status = categoriesResponse.status;
        throw error;
      }
  
      const postsJSON = await postsResponse.json();
      const categoriesR = await categoriesResponse.json();
  
      return {
        pagesR: postsJSON.metadata || postsJSON[0],
        postsR: postsJSON.posts || postsJSON[1],
        categoriesR,
      };
    } catch (error) {
      console.log(error);
      throw error;
    }
  };
//HOMEPAGE
  export const postsLoader = async ({ request }) => {
    const url = new URL(request.url);
    const currentQuery = url.search;
    return fetchPostsAndCategories("", currentQuery);
  };
//MY POSTS
  export const myPostsLoader = async ({ request }) => {
    const url = new URL(request.url);
    const currentQuery = url.search;
    return fetchPostsAndCategories("/post/get_my_posts/", currentQuery);
  };
//LIKED POSTS
  export const likedPostsLoader = async ({ request }) => {
    const url = new URL(request.url);
    const currentQuery = url.search;
    return fetchPostsAndCategories("/post/get_my_liked_posts/", currentQuery);
  };
