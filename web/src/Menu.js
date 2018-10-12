import React, { createElement } from 'react';
import { connect } from 'react-redux';
import compose from 'recompose/compose';
import { translate, MenuItemLink, Responsive, getResources } from 'react-admin';
import { withRouter } from 'react-router-dom';


const styles = {
    main: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'flex-start',
        height: '100%',
    },
};

const Menu = ({ onMenuClick, resources, translate, logout }) => (
    <div style={styles.main}>
        {resources.filter(r => r.icon).map(resource => (
            <MenuItemLink
                key={resource.name}
                to={`/${resource.name}`}
                primaryText={translate(`resources.${resource.name}.name`, {
                    smart_count: 2,
                })}
                leftIcon={createElement(resource.icon)}
                onClick={onMenuClick}
            />
        ))}
        <Responsive xsmall={logout} medium={null} />
    </div>
);

const mapStateToProps = state => ({
    resources: getResources(state),
});

const enhance = compose(
    withRouter,
    connect(mapStateToProps),
    translate
);

export default enhance(Menu);
