import React from 'react';
import { hot } from 'react-hot-loader';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';

class LoginCallback extends React.Component {

    constructor(props) {
        super(props);
    }
    componentDidMount() {
        let token = this.props.location.search.slice(1);
        this.props.setToken(token);
    }

    render() {
        return <Redirect to="/" />;
    }
}

const mapState = state => ({});
const mapDispatch = dispatch => ({
    setToken: token => dispatch({ type: 'JWT_RECEIVED', token: token }) 
});
export default hot(module)(connect(mapState, mapDispatch)(LoginCallback))