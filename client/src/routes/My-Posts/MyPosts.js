import { useLoaderData, useLocation } from 'react-router';
import PostInfo from '../../components/PostInfo/PostInfo';
import Categories from '../../components/Categories/Categories';
import { useState } from 'react';
import { Pagination } from '@mui/material';
import { useEffect } from 'react';
import { useNavigate } from 'react-router';

const MyPostsPage = () => {
  const {pagesR, postsR, categoriesR} = useLoaderData();
  const [pages, setPages] = useState(pagesR);
  const [posts, setPosts] = useState(postsR);
  const [categories, setCategories] = useState(categoriesR);
  
  const navigate = useNavigate();
  const location = useLocation();
  const query = new URLSearchParams(location.search);
  const category = query.get('category');
  
  useEffect(() => {
    setPages(pagesR)
    setPosts(postsR)
    setCategories(categoriesR)
  },[pagesR, postsR, categoriesR])

  const handlePageClick = async (e, value) => {
    e.preventDefault();
    if (value) {
      // Update the URL with the new page number
      const newQuery = new URLSearchParams(query);
      newQuery.set('page', value);
      if (category) {
        newQuery.set('category', category);
      }
      navigate({ search: newQuery.toString() }); // navigate to update state values with useEffect
    }
  };

  return (
    <>
      <Categories categories={categories} query={category}/>

      {category===null && posts && posts.map((post) => {
        return (<PostInfo key={post.Id} avatar={post.AuthorAvatar} postid={post.Id} title={post.Title} content={post.Content} date={post.CreationDate} author={post.Author}/>) 
      })}
  
      {category !== null && posts && posts.filter((post) => post.Category === category).map((post) => {
        return (<PostInfo key={post.Id} avatar={post.AuthorAvatar} postid={post.Id} title={post.Title} content={post.Content} date={post.CreationDate} author={post.Author}/>)
      })}

      {posts && <Pagination 
          count={pages.Pages}
          variant='outlined' 
          color="primary" 
          size='large' 
          onChange={(e, value) => handlePageClick(e, value)} 
          sx={{
            position:'fixed',
            bottom: 20, 
            left: "50%", 
            transform: "translateX(-50%)",
            '& .MuiPaginationItem-root': {
              backgroundColor: 'white',
              color: 'black',
              borderColor: 'rgba(0, 0, 0, 0.23)',
            },
            '& .MuiPaginationItem-root:hover': {
              borderColor: 'rgba(0, 0, 0, 0.23)',
            }
          }}
        />}
    </>
  )
}
export default MyPostsPage;