﻿#region Copyright (c) 2009-2016 Misakai Ltd.
/*************************************************************************
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU Affero General Public License as
* published by the Free Software Foundation, either version 3 of the
* License, or(at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
*  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the
* GNU Affero General Public License for more details.
*
* You should have received a copy of the GNU Affero General Public License
* along with this program.If not, see<http://www.gnu.org/licenses/>.
*************************************************************************/
#endregion Copyright (c) 2009-2016 Misakai Ltd.

using System;
using Newtonsoft.Json;

namespace Emitter.Network
{
    /// <summary>
    /// Represents a response for an API request.
    /// </summary>
    internal class EmitterResponse
    {
        /// <summary>
        /// Constructs a default response.
        /// </summary>
        public EmitterResponse()
        {
            this.Status = 200;
        }

        /// <summary>
        /// Constructs a default response.
        /// </summary>
        /// <param name="status">The status to set.</param>
        public EmitterResponse(int status)
        {
            this.Status = status;
        }

        /// <summary>
        /// The error status code for this response.
        /// </summary>
        [JsonProperty("status")]
        public int Status;
    }
}