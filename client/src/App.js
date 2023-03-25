import {
    createBrowserRouter, 
    createRoutesFromElements,
    Route, 
    RouterProvider,
  } from 'react-router-dom'  
  // pages
  import HomePage from './routes/Home-Page/HomePage'
  import SignInPage from './routes/Sign-In-Page/SignInPage'
  import SignUpPage from './routes/Sign-Up-Page/SignUpPage'
  import PostPage from './routes/Post-Page/PostPage'
  import CreatePost from './routes/Create-Post/CreatePost'
  import ProfilePage from './routes/Profile-Page/ProfilePage'
  import MyPostsPage from './routes/My-Posts/MyPosts'
  import LikedPosts from './routes/Liked-Posts/LikedPosts'
  
  // layouts
  import RootLayout from './layouts/RootLayer/RootLayout' 
  // loaders
  import { postsLoader, postLoader, checkAuth, myPostsLoader, likedPostsLoader, profileData, createPostLoader } from './api/Loaders'

  // errorElement
  import { ErrorPage } from './routes/Error-Page/ErrorPage'
  import { PageNotFound } from './routes/Error-Page/PageNotFound'
  
  const router = createBrowserRouter(
    createRoutesFromElements(

      <Route element={<RootLayout />} id="root" loader={checkAuth} errorElement={<ErrorPage/>}>
          <Route path="/" >
              <Route index 
                element={<HomePage />}
                loader={postsLoader}
              />
              <Route 
                path="createpost" 
                element={<CreatePost />}
                loader={createPostLoader}
              />
              <Route 
                path="posts"
                element={<PostPage />}
                loader={postLoader}  
              />
              <Route 
                path="my-posts"
                element={<MyPostsPage />}
                loader={myPostsLoader}  
              />
              <Route 
                path="liked-posts"
                element={<LikedPosts />}
                loader={likedPostsLoader}  
              />
              
              <Route 
                path="profile"
                element={<ProfilePage />}
                loader={profileData}  
              />              
          </Route>
        <Route path="signin" element={<SignInPage/>} />
        <Route path="signup" element={<SignUpPage/>}  />
        <Route path="*" element={<PageNotFound/>} />
      </Route>
    )
  )
  
  function App() {
    return (
      <RouterProvider router={router} />
    );
  }
  
  export default App