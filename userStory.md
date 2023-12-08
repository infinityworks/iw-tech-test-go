# User Story
As a user of the food ratings app I want to see food hygiene ratings as percentages in a chosen local authority so that I can understand the profile of establishments in that authority

# Elaboration
The existing application returns only one of two preset results. Expand the functionality to retrieve actual ratings from the Food Ratings API, calculate percentages and return them in order to the UI.

The Food Ratings API documentation can be found here: http://api.ratings.food.gov.uk/help

In England the Food Standards Agency rates establishments using the FHRS rating scheme. This uses a star rating from 5 to 0, or "Exempt".

In Scotland the Food Standards Agency rates establishments using the FHIS rating scheme. This uses a the values "Pass and Eat Safe", "Pass" and "Improvement Required".

For the purposes of this implementation the ratings schemes shall be considered fixed and unlikely to change.

# Return order
FHRS ratings should be returned in order, as in this example
<table>
    <thead>
        <tr>
            <th>Rating</th>
            <th>Percentage</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>5</td>
            <td>50%</td>
        </tr>
        <tr>
            <td>4</td>
            <td>0%</td>
        </tr>
        <tr>
            <td>3</td>
            <td>0%</td>
        </tr>
        <tr>
            <td>2</td>
            <td>0%</td>
        </tr>
        <tr>
            <td>1</td>
            <td>20%</td>
        </tr>
        <tr>
            <td>Exempt</td>
            <td>30%</td>
        </tr>
    </tbody>
</table>

FHIS ratings should be returned in order, as in this example
<table>
    <thead>
        <tr>
            <th>Rating</th>
            <th>Percentage</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Pass and Eat Safe</td>
            <td>50%</td>
        </tr>
        <tr>
            <td>Pass</td>
            <td>15%</td>
        </tr>
        <tr>
            <td>Improvement Required</td>
            <td>35%</td>
        </tr>
    </tbody>
</table>

# Acceptance Criteria
On selecting a local authority the rating values and percentage shall be displayed.

All possible values for the relevant scheme type shall be displayed in the correct order.