// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'

import Icon from '@ttn-lw/components/icon'

import Message from '@ttn-lw/lib/components/message'
import ErrorMessage from '@ttn-lw/lib/components/error-message'

import PropTypes from '@ttn-lw/lib/prop-types'
import { getEntityId } from '@ttn-lw/lib/selectors/id'
import { getBackendErrorRootCause } from '@ttn-lw/lib/errors/utils'

import Event from '../..'

import style from './error.styl'

const ErrorEvent = function({ className, event, expandedClassName, overviewClassName, widget }) {
  const entityId = getEntityId(event.identifiers[0])
  const icon = <Icon icon="error" className={style.icon} />
  // Transform error details to error-like structure.
  const rootCause = getBackendErrorRootCause({ details: [event.data] })

  const content = (
    <span className={style.content}>
      <Message content={{ id: `event:${event.name}` }} />
      {!widget && <ErrorMessage className={style.errorDescription} content={rootCause} />}
    </span>
  )

  return (
    <Event
      className={className}
      overviewClassName={overviewClassName}
      expandedClassName={expandedClassName}
      icon={icon}
      time={event.time}
      emitter={entityId}
      content={content}
      widget={widget}
      data={event.data}
    />
  )
}

ErrorEvent.propTypes = {
  className: PropTypes.string,
  event: PropTypes.event.isRequired,
  expandedClassName: PropTypes.string,
  overviewClassName: PropTypes.string,
  widget: PropTypes.bool,
}

ErrorEvent.defaultProps = {
  className: undefined,
  expandedClassName: undefined,
  overviewClassName: undefined,
  widget: false,
}

export default ErrorEvent
