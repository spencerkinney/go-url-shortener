import { Link } from "react-router-dom";
import data from "../../apiDocs.json";
import "./Docs.css";

const Docs = () => (
  <div id="container">
    <div>
      <span className="bold">Server URL: &nbsp;</span>
      {`${data.host}${data.port ? `:${data.port}` : ''}`}
    </div>

    { data.requests.map((request) => (
      <div id="requestContainer" key={request.endpoint}>
        <div id="requestTypeContainer">
          <span id='endpoint' className="bold">{request.endpoint}</span>
          <span
            id="requestType"
            className={`${request.requestType.toLowerCase()}Request bold`}
          >
            {request.requestType}
          </span>
        </div>
        <div id="requestContent">
          <div id="description">{request.description}</div>
          { request.accepts.length > 0 && (
            <div id="accepts">
              <span className='bold'>Accepts: &nbsp;</span>
              {request.accepts.join(" ")}
            </div>
          )}
          { request.produces.length > 0 && (
            <div id="produces">
              <span className='bold'>Produces: &nbsp;</span> 
              {request.accepts.join(" ")}
            </div>
          )}
          <div className="preWrap">
            <span className="bold">Request body: </span>
            <div className="body">
              {request.requestBody ?? 'No request body'}
            </div>
          </div>

          <div className="bold section">Responses: </div>
          { request.responses.map((response) => (
            <div key={response.statusCode}>
              <div>
                <span className="bold">{response.statusCode}</span> -&nbsp;
                {response.description}
              </div>
              <div className="preWrap body">
                {response.responseBody.join("\nor\n")}
              </div>
            </div>
          ))}
        </div>
      </div>
    ))}

    <Link id='link' to='/'>
      View Shortener
    </Link>
  </div>
);

export default Docs;