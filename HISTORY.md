# v0.2.3 - 07.14.2013
* Rename manages SetEngine method to Engine.
* Make managers engine member private to enforce using Engine method to modify it.
* Remove near useless type switch in favor of more recognizable type if.

# v0.2.2 - 07.11.2013
* Implement Having and Group By for Select Managers.
* Consistent naming convention for receivers.
* Consistent returns from factory methods.

# v0.2.1 - 07.10.2013
* Select Statement ordering with Ascending and Descending nodes.

# v0.1.2 - 07.09.2013
* Proper tests for Nodes
* Proper tests for ToSqlVisitor
* Visitor's Accept now returns result, error after recovering from a potential panic from the visitors Visit method.
* Apply check to all SelectCoreNodes's that belong to a SelectStatementNode to ensure a minimum of a STAR is present.
* Update tree managers to return the same string, error combination the VisitorInterface returns.
* Update README to reflect changes to VisitorInerfaces's Accept method and its' affects on tree managers.
* Remove all panics to allow for visitor to return any errors.
* Add DEBUG variable to visitor for debugging information.
* Add tests and factory methods for all managers.

# v0.1.1 - 07.08.2013
* Overhaul nodes package to allow for "factory" methods.
* Add Nodes and visitors for SQL Functions Average, Sum, Minimum, Maximum and Count.
* More documentation.
* Remove check that looked at number of projections in a SelectCoreNode, and placed job in SelectManager's ToSql receiver method.
