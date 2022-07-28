import * as React from 'react';
import { styled, alpha } from '@mui/material/styles';
import AppBar from '@mui/material/AppBar';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Button from '@mui/material/Button';
import InputBase from '@mui/material/InputBase';
import logo from '../../assets/logopng.png';
import SearchIcon from '@mui/icons-material/Search';
import Modal from '@mui/material/Modal';
import TextField from '@mui/material/TextField';
import { useNavigate } from "react-router-dom";

const style = {
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 1200,
    bgcolor: 'background.paper',
    border: '2px solid #000',
    boxShadow: 24,
    p: 4,
};

const Search = styled('div')(({ theme }) => ({
    position: 'relative',
    borderRadius: theme.shape.borderRadius,
    backgroundColor: alpha(theme.palette.common.white, 0.15),
    '&:hover': {
        backgroundColor: alpha(theme.palette.common.white, 0.25),
    },
    marginRight: theme.spacing(2),
    marginLeft: 0,
    width: '100%',
    [theme.breakpoints.up('sm')]: {
        marginLeft: theme.spacing(3),
        width: 'auto',
    },
}));

const SearchIconWrapper = styled('div')(({ theme }) => ({
    padding: theme.spacing(0, 2),
    height: '100%',
    position: 'absolute',
    pointerEvents: 'none',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
}));

const StyledInputBase = styled(InputBase)(({ theme }) => ({
    color: 'inherit',
    '& .MuiInputBase-input': {
        padding: theme.spacing(1, 1, 1, 0),
        // vertical padding + font size from searchIcon
        paddingLeft: `calc(1em + ${theme.spacing(4)})`,
        transition: theme.transitions.create('width'),
        width: '100%',
        [theme.breakpoints.up('md')]: {
            width: '50ch',
        },
    },
}));

export default function AppBarComponent() {
    const [open, setOpen] = React.useState(false);
    const navigate = useNavigate();
    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);

    const onClickPerfil = () => {
        navigate("/myProfile");
    };


    const onClickHome = () => {
        navigate("/home");
    };

    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar position="static" style={{ background: '#2292A4' }}>
                <Toolbar>
                    <img onClick={onClickHome} style={{ width: '200px', height: '100px', marginLeft: "135px" }} src={logo} />
                    <Search style={{ marginLeft: "95px" }}>
                        <SearchIconWrapper>
                            <SearchIcon />
                        </SearchIconWrapper>
                        <StyledInputBase
                            placeholder="Search…"
                            inputProps={{ 'aria-label': 'search' }}
                        />
                    </Search>
                    <Box sx={{ flexGrow: 1 }} />
                    <Box style={{ marginRight: "25px" }} sx={{ display: { md: 'flex' } }}>
                        <Button onClick={onClickPerfil} color="inherit">Mi Perfil</Button>
                    </Box>
                    <Box style={{ marginRight: "95px" }} sx={{ display: { md: 'flex' } }}>
                        <Button onClick={handleOpen} style={{ background: "rgba(102, 229, 250, 0.44)" }} variant="contained">Agregar Oferta</Button>
                    </Box>
                </Toolbar>
            </AppBar>
            <Modal
                open={open}
                onClose={handleClose}
                aria-labelledby="modal-modal-title"
                aria-describedby="modal-modal-description"
            >
                <Box sx={style}>
                    <Grid container direction='row'>
                        <Grid xs={6} style={{ paddingRight: '30px' }} >
                            <h1>Producto </h1>
                            <TextField fullWidth id="outlined-basic" label="Nombre del Producto" variant="outlined" />
                            <Button onClick={handleClose} style={{ background: "#2292A4", marginTop: '90px', marginLeft: '190px' }} variant="contained">Agregar</Button>

                        </Grid>
                        <Grid xs={6} direction='row'>
                            <h1>Oferta </h1>
                            <TextField fullWidth id="outlined-basic" label="Descripcion" variant="outlined" />
                            <TextField style={{ marginTop: '30px' }} fullWidth id="outlined-basic" label="Precio" variant="outlined" />
                            <TextField style={{ marginTop: '30px' }} fullWidth id="outlined-basic" label="Link" variant="outlined" />
                        </Grid>
                    </Grid>
                </Box>
            </Modal>
        </Box>
    );
}
