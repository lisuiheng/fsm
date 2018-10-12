import React from 'react';
import {
    List,
    Create,
    Edit,
    TabbedForm,
    FormTab,
    TextInput,
    required,
    Show,
    SimpleShowLayout,
    Responsive,
    Datagrid,
    SelectField,
    TextField,
    ReferenceField,
    SelectInput,
} from 'react-admin';
import withStyles from '@material-ui/core/styles/withStyles';
import Icon from '@material-ui/icons/Collections';

export const UserIcon = Icon;
export const UserResourceName = 'users';

const ROLE_NO = 1;
const ROLE_USER = 100;
const ROLE_ADMIN = 200;

const roleChoices = [
    { id: ROLE_NO, name: 'user.choices.role.no' },
    { id: ROLE_USER, name: 'user.choices.role.user' },
    { id: ROLE_ADMIN, name: 'user.choices.role.admin' },
];

export const UserList = props => (
    <List {...props} perPage={20} sort={{ field: 'id', order: 'ASC' }}>
        <Responsive
            medium={
                <Datagrid>
                    <ReferenceField
                        label="resources.users.fields.name"
                        source="id"
                        reference="users"
                        linkType="show"
                    >
                        <TextField source="name" />
                    </ReferenceField>
                    <SelectField
                        source="role"
                        choices={roleChoices}
                        translateChoice={true}
                    />
                </Datagrid>
            }
        />
    </List>
);

const createStyles = {
    name: { width: '5em' },
};

export const UserCreate = withStyles(createStyles)(({ classes, ...props }) => (
    <Create {...props}>
        <TabbedForm>
            <FormTab label="resources.users.tabs.detail">
                <TextInput
                    source="name"
                    validate={required()}
                    options={{ fullWidth: true }}
                />
                <TextInput source="username" options={{ fullWidth: true }} />
                <TextInput source="password" options={{ fullWidth: true }} />
                <SelectInput
                    source="role"
                    choices={roleChoices}
                    translateChoice={true}
                />
            </FormTab>
        </TabbedForm>
    </Create>
));

const UserTitle = ({ record }) => <span>Poster #{record.reference}</span>;

const editStyles = {
    ...createStyles,
    comment: {
        maxWidth: '20em',
        overflow: 'hidden',
        textOverflow: 'ellipsis',
        whiteSpace: 'nowrap',
    },
};

export const UserEdit = withStyles(editStyles)(({ classes, ...props }) => (
    <Edit {...props} title={<UserTitle />}>
        <TabbedForm>
            <FormTab label="resources.users.tabs.detail">
                <TextInput source="name" options={{ fullWidth: true }} />
                <TextInput source="username" options={{ fullWidth: true }} />
                <TextInput source="password" options={{ fullWidth: true }} />
                <SelectInput
                    source="role"
                    choices={roleChoices}
                    translateChoice={true}
                />
            </FormTab>
        </TabbedForm>
    </Edit>
));

export const UserShow = props => (
    <Show {...props}>
        <SimpleShowLayout>
            <TextField source="name" />
            <TextField source="username" />
            <SelectField
                source="role"
                choices={roleChoices}
                translateChoice={true}
            />
        </SimpleShowLayout>
    </Show>
);
