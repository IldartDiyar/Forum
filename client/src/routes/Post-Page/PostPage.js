import { useLoaderData } from 'react-router';
import Post from '../../components/Post/Post';
import CommentComponent from '../../components/Comment/Comment';
import { Box, Typography, Grid } from '@mui/material';

const PostPage = () => {
const searchParams = new URLSearchParams(window.location.search);
const id = searchParams.get('id');
const { post, comments } = useLoaderData();

return (
<Box
  sx={{
  display: 'flex',
  flexDirection: 'column',
  }}
>
  {post.ImageData && (
  <Box
    component="img"
    src={post.ImageData}
    alt="Italian Trulli"
    sx={{ maxWidth: '100%' }}
  />
  )}
  <Post
      key={post.PostId}
      query={id}
      postid={post.PostId}
      title={post.Title}
      content={post.Content}
      date={post.CreationDate}
      author={post.Author}
      countLike={post.CountLike}
      countDislike={post.CountDislike}
      isLiked={post.Likeisset}
      isDisliked={post.Dislikeisset}
    />
  <Typography
    variant="h4"
    fontFamily="Dosis"
    sx={{
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    pt: 4,
    pb: 4,
    background: 'radial-gradient(farthest-corner at 50% 50%, #00C9A7, #845EC2) border-box',
    WebkitTextFillColor: 'transparent',
    WebkitBackgroundClip: 'text',
    }}
  >
  COMMENTS
  </Typography>
  <Grid container spacing={2} sx={{ width: '100%' }}>
  {comments !== null &&
  comments.map((comment) => {
    return (
    <Grid item xs={12} key={comment.Id}>
    <CommentComponent
                  commentID={comment.Id}
                  date={comment.CreationDate}
                  author={comment.Author}
                  content={comment.Body}
                  countLike={comment.CountLike}
                  countDislike={comment.CountDislike}
                  isLiked={comment.Likeisset}
                  isDisliked={comment.Dislikeisset}
                />
    </Grid>
    );
  })}
  </Grid>
</Box>
);
};
export default PostPage;