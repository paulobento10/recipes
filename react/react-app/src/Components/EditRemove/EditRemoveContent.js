import React, { Component } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { withStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import { Table, Divider, Tag } from 'antd';

const columns = [
  {
    title: 'Name',
    dataIndex: 'name',
    key: 'name',
    render: text => <a>{text}</a>,
  },
  {
    title: 'Age',
    dataIndex: 'age',
    key: 'age',
  },
  {
    title: 'Address',
    dataIndex: 'address',
    key: 'address',
  },
  {
    title: 'Action',
    key: 'action',
    render: (text, record) => (
      <span>
        <a>Edit</a>
        <Divider type="vertical" />
        <a>Delete</a>
      </span>
    ),
  },
];

const data = [
  {
    key: '1',
    name: 'John Brown',
    age: 32,
    address: 'New York No. 1 Lake Park',
  },
  {
    key: '2',
    name: 'Jim Green',
    age: 42,
    address: 'London No. 1 Lake Park',
  },
  {
    key: '3',
    name: 'Joe Black',
    age: 32,
    address: 'Sidney No. 1 Lake Park',
  },
];

function createData(name, code, population, size) {
  const density = population / size;
  return { name, code, population, size, density };
}

class EditRemoveContent extends Component {

  constructor(props){
    super(props);
    this.state = {
        page: 0,
        rowsPerPage: 10,
    };
    this.handleChangePage=this.handleChangePage.bind(this);
    this.handleChangeRowsPerPage=this.handleChangeRowsPerPage.bind(this);
}


    handleChangePage = (event, newPage) => {
        this.setState({page: newPage})
    };

    handleChangeRowsPerPage = event => {
        this.setState({rowsPerPage: +event.target.value})
        this.setState({page: 0})
    };

    render() {
    const { classes } = this.props;

    return (
        <Paper>
          <Table columns={columns} dataSource={data} />
        </Paper>
      );
    }
}

export default EditRemoveContent;