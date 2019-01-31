import * as React from 'react';
import { hot } from 'react-hot-loader/root';
import { Redirect } from 'react-router-dom';
import { connect } from 'react-redux';

class Props {
    setToken = (token) => {};
    location: any;
}
class LoginCallback extends React.Component<Props> {

    constructor(props: any) {
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
export default connect(mapState, mapDispatch)(hot(LoginCallback))