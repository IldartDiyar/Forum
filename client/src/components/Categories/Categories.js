import Stack from '@mui/material/Stack';
import Button from '@mui/material/Button';
import { useLocation, useNavigate } from 'react-router-dom';

export default function Categories(props) {
  const location = useLocation();
  const navigate = useNavigate();
  const searchParams = new URLSearchParams(location.search);
  const category = searchParams.get('category');
  
  // const page = searchParams.get('page') || 1;

  const handleCategoryClick = async (categoryName) => {
    const updatedSearchParams = new URLSearchParams();
    updatedSearchParams.set('category', categoryName);
    // updatedSearchParams.set('page', page);
    const updatedSearch = updatedSearchParams.toString();
    navigate({ pathname: location.pathname, search: updatedSearch });
  };

  return (
    <Stack
      direction="row"
      justifyContent="space-between"
      alignItems="center"
      sx={{
        position: "sticky",
        top: 0,
        zIndex: 3,
        height: "5vh",
      }}
    >
      {props.categories.map((elem) => {
        return (
          <Button
            onClick={() => handleCategoryClick(elem.categoryName)}
            key={elem.categoryId}
            sx={{
              height:'100%',
              width: '100%',
              fontFamily: 'Bebas Neue',
              fontSize: 20,
              borderRadius:0,
              backgroundColor: elem.categoryName === category ? '#62a1bd' : 'rgb(106, 113, 184)',
              color: 'white',
              border: '2px solid transparent',
              borderColor: 'black',
              '&:hover': {
                backgroundColor: '#b04e5e',
                borderColor: 'white',
              },
            }}
          >
            {elem.categoryName}
          </Button>
        );
      })}
    </Stack>
  );
}