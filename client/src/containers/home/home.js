import React, { useEffect } from 'react'


import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import MenuList from '@material-ui/core/MenuList';
import ClickAwayListener from '@material-ui/core/ClickAwayListener';
import Grow from '@material-ui/core/Grow';
import Paper from '@material-ui/core/Paper';
import Popper from '@material-ui/core/Popper';

import useStyles from './styles'

export default function Home(props) {
    const classes = useStyles();
    const [open, setOpen] = React.useState(false);
    const anchorRef = React.useRef(null);

    function handleToggle() {
        setOpen(prevOpen => !prevOpen);
    }

    function handleClose(event) {
        if (anchorRef.current && anchorRef.current.contains(event.target)) {
            return;
        }

        setOpen(false);
    }

    useEffect(() => {
        if (localStorage.getItem('user') === null) {
            props.history.replace('/signin')
        }
    }, []);

    return (
        <div className={classes.root}>
            <AppBar position="static">
                <Toolbar>
                    <div className={classes.menu}>
                        <IconButton
                            edge="start"
                            className={classes.menuButton}
                            color="inherit" aria-label="menu"
                            ref={anchorRef}
                            aria-controls="menu-list-grow"
                            aria-haspopup="true"
                            onClick={handleToggle}
                        >
                            <MenuIcon />
                        </IconButton>

                        <Popper open={open} anchorEl={anchorRef.current} transition disablePortal>
                            {({ TransitionProps, placement }) => (
                                <Grow
                                    {...TransitionProps}
                                    style={{ transformOrigin: placement === 'bottom' ? 'center top' : 'center bottom' }}
                                >
                                    <Paper id="menu-list-grow">
                                        <ClickAwayListener onClickAway={handleClose}>
                                        <MenuList>
                                            <MenuItem onClick={handleClose}>Profile</MenuItem>
                                            <MenuItem onClick={handleClose}>My account</MenuItem>
                                            <MenuItem onClick={handleClose}>Logout</MenuItem>
                                        </MenuList>
                                        </ClickAwayListener>
                                    </Paper>
                                </Grow>
                            )}
                        </Popper>
                    </div>

                    <Button color="inherit" href="/">Home</Button>
                    <Button color="inherit" href="/about">About</Button>
                    <Button color="inherit" href="/signin">Login</Button>
                </Toolbar>
            </AppBar>
        </div>
    );
}