package jago

const MAIN_METHOD_NAME = "main"
const MAIN_METHOD_DESCRIPTOR = "([Ljava/lang/String;)V"

func Startup(initialClassName string, args... string)  {

	RegisterNatives()

	// initial a thread
	THREAD_MANAGER.NewThread("main", func(thread *Thread) {
		// welcome to the Java world
		// the Java journey starts here
		defer func() { // finally check uncaught exception
			r := recover()
			if r != nil {
				throwable, ok := r.(Reference)
				if ok {
					thread.handleUncaughtException(throwable)
				}
			}
		}()

		VM_invokeMethod("java/lang/System", "initializeSystemClass", "()V")

		// once system classes are initialized, set up the java.lang.Thread object
		threadConstructor := thread.threadObject.Class().GetConstructor("(Ljava/lang/String;)V")
		VM_invokeMethod0(threadConstructor, thread.threadObject, NewJavaLangString(thread.name))

		// As per jvm specification, initial main method needs to initialize initial class
		argsArr := NewArray("[Ljava/lang/String;", Int(len(args)))
		for i, arg := range args {
			argsArr.SetElement(Int(i), NewJavaLangString(arg))
		}
		VM_invokeMethod(initialClassName, MAIN_METHOD_NAME, MAIN_METHOD_DESCRIPTOR, argsArr)

	})


}
