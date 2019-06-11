import React, {Component} from 'react';

class Form extends Component {
    constructor(props) {
        super(props);

        this.initialState = {
            name: '',
            job: ''
        };

        this.state = this.initialState;
    }

    // 下面是对事件的处理，event 会被传递
    handleChange = event => {
        const {name, value} = event.target;

        this.setState({
            [name]: value
        });
    }

    //  onFormSubmit = (event) => {
    //     event.preventDefault();
    //
    //     this.props.handleSubmit(this.state);
    //     this.setState(this.initialState);
    // }

    submitForm = () => {
        // event.preventDefault() 会阻碍默认的操作
        // 参考：https://hashnode.com/post/why-do-you-write-eventpreventdefault-in-react-cjdznf1el0atom3wt831c2m9o

        event.preventDefault();
        this.props.handleSubmit(this.state);
        this.setState(this.initialState);
    }

    render() {
        // TODO 感觉像是 js 解包的语法糖
        const {name, job} = this.state;
        return (
            <form onSubmit={this.submitForm}>
                <label>name</label>
                <input
                    type="text"
                    name="name"
                    value={name}
                    onChange={this.handleChange}
                />
                <label>job</label>
                <input
                    type="text"
                    name="job"
                    value={job}
                    onChange={this.handleChange}
                />
                <button type="submit">
                    Submit
                </button>
            </form>
        );
    }
}

export default Form;