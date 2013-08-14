# v0.3.5 - 08.14.2013
* Introduce nodes, visitors, and managers for table creation and alteration.

# v0.3.4 - 08.13.2013
* Major package update to remove unused "tree" directory as it does not represent a package.
* Returning method on an InsertManager should only covert strings passed as arguments to Column nodes.
* Allow for columns within an INSERT statements to be SQL bindings.
* Replace engine field and Engine method on managers in favor of adapter and SetAdapter respectively.

# v0.3.3 - 08.01.2013
* Implement BindingNode for SQL parameterization.
* Implement BindingNode for PostgreSQL parameterization.

# v0.3.2 - 07.27.2013
* Update to Where method for SelectMananger - if provided expression is a string, convert it to a literal node.
* Update to Order method for SelectMananger - if provided expression is a string, convert it to a literal node.
* Update to Having method for SelectMananger - if provided expression is a string, convert it to a literal node.
* Update to Where method for SelectMananger - always group expression unless the exression is a group already.

# v0.3.1 - 07.25.2013
* Add Returning method for InsertManager's (geared toward Postgres).

# v0.3 - 07.18.2013
* Add initial MySQL visitor.
* Implement UNION, INERSECT, and EXCEPT for Select statments.

# v0.2.3 - 07.14.2013
* Rename managers SetEngine method to Engine.
* Make managers engine member private to enforce using Engine method to modify it.
* Remove near useless type switch in favor of more recognizable type if.

# v0.2.2 - 07.11.2013
* Implement Having and Group By for SelectManagers.
* Consistent naming convention for receivers.
* Consistent returns from factory methods.

# v0.2.1 - 07.10.2013
* Select Statement ordering with Ascending and Descending nodes.

# v0.1.2 - 07.09.2013
* Proper tests for Nodes.
* Proper tests for ToSqlVisitor.
* Visitor's Accept method now returns result, error after recovering from a potential panic from the visitors Visit method.
* Apply check to all SelectCoreNodes's that belong to a SelectStatementNode to ensure a minimum of a StarNode is present.
* Update tree managers to return the same string, error combination the VisitorInterface returns.
* Update README to reflect changes to VisitorInerfaces's Accept method and its affects on tree managers.
* Remove all panics to allow for visitor to return any errors.
* Add DEBUG variable to visitor for debugging information.
* Add tests and factory methods for all managers.

# v0.1.1 - 07.08.2013
* Overhaul nodes package to allow for "factory" methods.
* Add Nodes and visitors for SQL Functions Average, Sum, Minimum, Maximum and Count.
* More documentation.
* Remove check that ensure a minimum of a StarNode was found in a SelectCoreNode's projections, and placed job in SelectManager's ToSql receiver method.
